package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a made of multiple rows.
type Matrix [][]int

// Rows returns a slice of the matrix rows.
func (m *Matrix) Rows() [][]int {
	matrix := *m
	rows := make([][]int, len(matrix))
	for i := range matrix {
		rows[i] = make([]int, len(matrix[0]))
		copy(rows[i], matrix[i])
	}
	return rows
}

// Cols returns a slice of the matrix columns.
func (m *Matrix) Cols() [][]int {
	matrix := *m
	cols := make([][]int, len(matrix[0]))
	for _, row := range matrix {
		for i, v := range row {
			cols[i] = append(cols[i], v)
		}
	}
	return cols
}

// Set updates a matrix element. It returns false if the index
// does not exist.
func (m *Matrix) Set(row, col, val int) bool {
	matrix := *m
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return false
	}
	matrix[row][col] = val
	return true
}

// New contructs a matrix from the given input. It returns an error
// if the input contains a bad value or inconsistent dimensions.
func New(input string) (*Matrix, error) {
	var m Matrix
	inputRows := strings.Split(input, "\n")
	rowLength := 0
	for _, ir := range inputRows {
		row := []int{}
		for _, i := range strings.Split(strings.TrimSpace(ir), " ") {
			value, err := strconv.Atoi(i)
			if err != nil {
				return &Matrix{}, errors.New("input contains bad value")
			}
			row = append(row, value)
		}
		if len(m) == 0 {
			rowLength = len(row)
		}
		if len(row) != rowLength {
			return &Matrix{}, errors.New("inconsistent dimensions")
		}
		m = append(m, row)
	}
	return &m, nil
}
