package lapjv_test

import (
	"testing"

	"github.com/heetch/lapjv"
	"github.com/stretchr/testify/assert"
)

func TestLapjvSampleMatrix(t *testing.T) {
	matrix := newMatrix(4, 4)
	matrix[0] = []int{5, 2, 9, 2}
	matrix[1] = []int{6, 4, 6, 2}
	matrix[2] = []int{2, 4, 5, 1}
	matrix[3] = []int{2, 4, 5, 1}

	m, n := 4, 4
	u := make([]int, n)
	v := make([]int, n)
	colsol := make([]int, n)
	rowsol := make([]int, n)
	lapcost := lapjv.Lapjv(m, matrix, rowsol, colsol, u, v)
	assert.Equal(t, []int{1, 2, 3, 0}, rowsol)
	assert.Equal(t, []int{3, 0, 1, 2}, colsol)
	assert.Equal(t, 11, lapcost)
}

func TestLapjvDiagonalMatrix(t *testing.T) {
	matrix := CreateDiagonalMatrix(10, 10)
	m, n := 10, 10
	u := make([]int, n)
	v := make([]int, n)
	colsol := make([]int, n)
	rowsol := make([]int, n)
	lapcost := lapjv.Lapjv(m, matrix, rowsol, colsol, u, v)
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, rowsol)
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, colsol)
	assert.Equal(t, 120, lapcost)
}

func CreateDiagonalMatrix(m, n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)

		for j := 0; j < n; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}
