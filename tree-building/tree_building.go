package tree

import (
	"errors"
	"sort"
)

// Record contains an ID and the ID of the parent.
type Record struct {
	ID     int
	Parent int
}

// Node contains an ID and a slice of pointers to it's children.
type Node struct {
	ID       int
	Children []*Node
}

// Build returns the root of a tree recreated from a slice of records.
// An error is returned if the IDs are not contiguous, increasing from 0,
// or a non-root ID has a parent ID equal to or smaller than itself.
func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node, len(records))
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.ID < r.Parent || (r.ID != 0 && r.ID == r.Parent) {
			return &Node{}, errors.New("invalid records")
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID > 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
