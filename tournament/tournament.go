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
	wins    int
	draws   int
	losses  int
	matches int
	points  int
}

// Tally writes the given match results as a formatted table, returning
// an error if it recieves bad data.
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	tally := map[string]record{}
	teams := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// ignore comments and newlines
		if line == "" || line[0] == '#' {
			continue
		}
		tokens := strings.Split(line, ";")
		if len(tokens) != 3 {
			return errors.New("bad results format")
		}
		team1Record, team1Ok := tally[tokens[0]]
		team2Record, team2Ok := tally[tokens[1]]
		if !team1Ok {
			teams = append(teams, tokens[0])
		}
		if !team2Ok {
			teams = append(teams, tokens[1])
		}
		team1Record.matches++
		team2Record.matches++
		switch tokens[2] {
		case "win":
			team1Record.wins++
			team1Record.points += 3
			team2Record.losses++
		case "loss":
			team1Record.losses++
			team2Record.wins++
			team2Record.points += 3
		case "draw":
			team1Record.draws++
			team1Record.points++
			team2Record.draws++
			team2Record.points++
		default:
			return errors.New("bad match result")
		}
		tally[tokens[0]] = team1Record
		tally[tokens[1]] = team2Record
	}
	// sort by wins, break ties alphabetically
	sort.SliceStable(teams, func(i, j int) bool {
		if tally[teams[i]].points == tally[teams[j]].points {
			name1 := strings.ToLower(teams[i])
			name2 := strings.ToLower(teams[j])
			for k := 0; k < len(name1); k++ {
				if name1[k] < name2[k] {
					return true
				}
				if name1[k] > name2[k] {
					return false
				}
			}
			return true
		}
		return tally[teams[i]].points > tally[teams[j]].points
	})
	// write formatted table
	fmt.Fprintf(w, "Team%s| MP |  W |  D |  L |  P\n", strings.Repeat(" ", 27))
	for _, t := range teams {
		r := tally[t]
		fmt.Fprintf(w, "%v%s|  %v |  %v |  %v |  %v |  %v\n",
			t, strings.Repeat(" ", 31-len(t)), r.matches, r.wins, r.draws,
			r.losses, r.points)
	}
	return nil
}
