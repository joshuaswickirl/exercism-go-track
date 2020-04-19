package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a made of multiple rows.
type Matrix struct {
	data       [][]int
	rows, cols int
}

// Rows returns a slice of the matrix rows.
func (m *Matrix) Rows() [][]int {
	c := make([][]int, m.rows)
	copy(c, m.data)
	return c
}

// Cols returns a slice of the matrix columns.
func (m *Matrix) Cols() [][]int {
	c := make([][]int, m.cols)
	for _, row := range m.data {
		for j, e := range row {
			c[j] = append(c[j], e)
		}
	}
	return c
}

// Set updates a matrix element. It returns false if the index
// does not exist.
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return false
	}
	m.data[row][col] = val
	return true
}

// New contructs a matrix from the given input. It returns an error
// if the input contains a bad value or inconsistent dimensions.
func New(input string) (*Matrix, error) {
	var m Matrix
	inputRows := strings.Split(input, "\n")
	for _, ir := range inputRows {
		row := []int{}
		for _, i := range strings.Split(strings.TrimSpace(ir), " ") {
			e, err := strconv.Atoi(i)
			if err != nil {
				return &Matrix{}, errors.New("input contains bad value")
			}
			row = append(row, e)
		}
		if m.cols == 0 {
			m.cols = len(row)
			m.rows = len(inputRows)
		}
		if len(row) != m.cols {
			return &Matrix{}, errors.New("inconsistent dimensions")
		}
		m.data = append(m.data, row)
	}
	return &m, nil
}
