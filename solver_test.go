package lapjv

import "testing"

func generateMatrix(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
		for j := 0; j < size; j++ {
			m[i][j] = i*j
		}
	}
	return m
}

func TestSmall(t *testing.T) {
	m := generateMatrix(10)
	s := MatrixSolver(m)

	if s.Cost != 120 {
		t.Errorf("Failure with small matrix solving : expected cost : 120 - res : %d", s.Cost)
	}
	if len(s.Rowsol) != 10 {
		t.Errorf("Failure with small matrix solving : expected len(rowsol) : 10 - res : %d", len(s.Rowsol))
	}
	if len(s.Colsol) != 10 {
		t.Errorf("Failure with small matrix solving : expected len(colsol) : 10 - res : %d", len(s.Colsol))
	}
	//We will check some values of the rowsol / colsol slices.

	if s.Colsol[2] != 7 {
		t.Errorf("Failure with small matrix solving : expected colsol[2] : 7 - res : %d", s.Colsol[2])
	}
	if s.Colsol[5] != 4 {
		t.Errorf("Failure with small matrix solving : expected colsol[5] : 4 - res : %d", s.Colsol[5])
	}
	if s.Rowsol[0] != 9 {
		t.Errorf("Failure with small matrix solving : expected rowsol[0] : 9 - res : %d", s.Rowsol[0])
	}
}