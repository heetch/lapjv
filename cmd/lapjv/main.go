package main

import (
	"fmt"

	"github.com/heetch/lapjv"
)

func main() {
	matrix := createWorstMatrix(10, 10)
	result := lapjv.Lapjv(matrix)
	fmt.Printf("Cost %d\n", result.Cost)
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
