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
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Prepare for building the tree by creating a map of IDs and their
	// *Node. Detect non-continous ID's by comparing the GaussianSum to
	// the actual sum of the ID's.
	nodeByID := map[int]*Node{}
	GaussianSum := (float64(len(records)) / 2) * (0 + float64(len(records)) - 1)
	sumOfIDs := 0.0
	for _, record := range records {
		if _, ok := nodeByID[record.ID]; ok {
			return &Node{}, errors.New("duplicate node")
		} else if record.ID == 0 && record.Parent != 0 {
			return &Node{}, errors.New("root node has parent")
		} else if record.ID <= record.Parent && record.ID != 0 {
			return &Node{}, errors.New("record has impossible child")
		}
		nodeByID[record.ID] = &Node{ID: record.ID}
		sumOfIDs += float64(record.ID)
	}
	if sumOfIDs != GaussianSum {
		return &Node{}, errors.New("non-continuous ID's")
	}

	// Build a tree and return the root.
	root, ok := nodeByID[0]
	if !ok {
		return &Node{}, errors.New("no root node")
	}
	for _, record := range records {
		if record.ID == 0 {
			continue
		}
		parent := nodeByID[record.Parent]
		parent.Children = append(parent.Children, nodeByID[record.ID])
		sort.SliceStable(parent.Children, func(i, j int) bool {
			return parent.Children[i].ID < parent.Children[j].ID
		})
	}
	return root, nil
}
