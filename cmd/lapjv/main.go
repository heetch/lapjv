package main

import (
	"fmt"

	"github.com/heetch/lapjv"
)

func main() {
	m, n := 1000, 1000
	u := make([]int, n)
	v := make([]int, n)
	colsol := make([]int, n)
	rowsol := make([]int, n)
	matrix := createWorstMatrix(m, n)

	lapcost := lapjv.Lapjv(m, matrix, rowsol, colsol, u, v)
	fmt.Printf("Cost %d\n", lapcost)
}

func createWorstMatrix(m, n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)

		for j := 0; j < n; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}
