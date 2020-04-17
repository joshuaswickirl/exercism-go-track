package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type record struct {
	wins   int
	draws  int
	losses int
}

func (r record) matches() int {
	return r.wins + r.draws + r.losses
}

func (r record) points() int {
	return 3*r.wins + r.draws
}

// Tally writes the given match results as a formatted table, returning
// an error if it recieves bad data.
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	tally := map[string]*record{}
	teams := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// ignore comments and newlines
		if line != "" && line[0] != '#' {
			tokens := strings.SplitN(line, ";", -1)
			if len(tokens) != 3 {
				return errors.New("bad results format")
			}
			team1, team2, result := tokens[0], tokens[1], tokens[2]
			if !strings.Contains(team1, " ") || !strings.Contains(team2, " ") {
				return errors.New("bad team name")
			}
			if tally[team1] == nil {
				tally[team1] = &record{}
				teams = append(teams, team1)
			}
			if tally[team2] == nil {
				tally[team2] = &record{}
				teams = append(teams, team2)
			}
			switch result {
			case "win":
				tally[team1].wins++
				tally[team2].losses++
			case "loss":
				tally[team1].losses++
				tally[team2].wins++
			case "draw":
				tally[team1].draws++
				tally[team2].draws++
			default:
				return errors.New("bad match result")
			}
		}
	}
	// sort by wins, break ties alphabetically, albeit only by first letter ;)
	sort.SliceStable(teams, func(i, j int) bool {
		if tally[teams[i]].points() == tally[teams[j]].points() {
			return byte(teams[i][0]) < byte(teams[j][0])
		}
		return tally[teams[i]].points() > tally[teams[j]].points()
	})
	// write formatted table
	output := []byte(fmt.Sprintf("Team%s| MP |  W |  D |  L |  P\n",
		strings.Repeat(" ", 27)))
	for _, t := range teams {
		r := tally[t]
		b := []byte(fmt.Sprintf("%v%s|  %v |  %v |  %v |  %v |  %v\n",
			t, strings.Repeat(" ", 31-len(t)), r.matches(), r.wins, r.draws,
			r.losses, r.points()))
		output = append(output, b...)
	}
	w.Write(output)
	return nil
}
