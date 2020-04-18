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
	team    string
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
		a, b := tally[tokens[0]], tally[tokens[1]]
		a.matches++
		b.matches++
		switch tokens[2] {
		case "win":
			a.wins++
			a.points += 3
			b.losses++
		case "loss":
			a.losses++
			b.wins++
			b.points += 3
		case "draw":
			a.draws++
			a.points++
			b.draws++
			b.points++
		default:
			return errors.New("bad match result")
		}
		tally[tokens[0]] = a
		tally[tokens[1]] = b
	}
	records := []record{}
	for t, r := range tally {
		r.team = t
		records = append(records, r)
	}
	// sort by wins, break ties alphabetically
	sort.Slice(records, func(i, j int) bool {
		if records[i].points == records[j].points {
			return strings.ToLower(records[i].team) < strings.ToLower(records[j].team)
		}
		return records[i].points > records[j].points
	})
	// write formatted table
	fmt.Fprintf(w, "Team%-27s| MP |  W |  D |  L |  P\n", "")
	for _, r := range records {
		fmt.Fprintf(w, "%-31s|  %v |  %v |  %v |  %v |  %v\n",
			r.team, r.matches, r.wins, r.draws, r.losses, r.points)
	}
	return nil
}
