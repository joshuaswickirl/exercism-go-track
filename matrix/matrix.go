package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a made of multiple rows.
type Matrix struct {
	data             [][]int
	numRows, numCols int
}

// Rows returns a slice of the matrix rows.
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.numRows)
	copy(rows, m.data)
	return rows
}

// Cols returns a slice of the matrix columns.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.numCols)
	for _, row := range m.data {
		for i, v := range row {
			cols[i] = append(cols[i], v)
		}
	}
	return cols
}

// Set updates a matrix element. It returns false if the index
// does not exist.
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.numRows || col < 0 || col >= m.numCols {
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
			value, err := strconv.Atoi(i)
			if err != nil {
				return &Matrix{}, errors.New("input contains bad value")
			}
			row = append(row, value)
		}
		if m.numCols == 0 {
			m.numCols = len(row)
			m.numRows = len(inputRows)
		}
		if len(row) != m.numCols {
			return &Matrix{}, errors.New("inconsistent dimensions")
		}
		m.data = append(m.data, row)
	}
	return &m, nil
}
