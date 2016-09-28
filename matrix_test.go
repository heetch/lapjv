package lapjv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareMatrixToSquare(t *testing.T) {
	squareMatrix := newMatrix(3, 3)
	squareMatrix[0] = []int{1, 2, 3}
	squareMatrix[1] = []int{1, 2, 3}
	squareMatrix[2] = []int{1, 2, 3}

	result := ToSquare(squareMatrix)

	assert.Equal(t, squareMatrix, result)
}

func TestVerticalMatrixToSquare(t *testing.T) {
	verticalMatrix := newMatrix(2, 1)
	verticalMatrix[0] = []int{1}
	verticalMatrix[1] = []int{1}

	expectedMatrix := newMatrix(2, 2)
	expectedMatrix[0] = []int{1, BIG}
	expectedMatrix[1] = []int{1, BIG}

	result := ToSquare(verticalMatrix)

	assert.Equal(t, expectedMatrix, result)
}

func TestHorizontalMatrixToSquare(t *testing.T) {
	horizontalMatrix := newMatrix(1, 2)
	horizontalMatrix[0] = []int{1, 2}

	expectedMatrix := newMatrix(2, 2)
	expectedMatrix[0] = []int{1, 2}
	expectedMatrix[1] = []int{BIG, BIG}

	result := ToSquare(horizontalMatrix)

	assert.Equal(t, expectedMatrix, result)
}

func newMatrix(x, y int) [][]int {
	rows := make([][]int, x)

	for i := range rows {
		rows[i] = make([]int, y)
	}

	return rows
}
