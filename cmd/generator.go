package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"

	"github.com/heetch/lapjv"
)

//MatrixGenerator will store info we will use to generate our matrix and store it.
//this type also contain the Matrix itself as a [][]int.
type MatrixGenerator struct {
	Size   int
	Type   FillType
	Matrix [][]int
}

//matrixGeneratorInitInteractive function will prompt the user for each config part to create a MatrixGenerator
func NewInteractiveMatrixGenerator() (*MatrixGenerator, error) {
	c := &MatrixGenerator{Type: Random}

	fmt.Print("Please enter the size of your matrix : ")
	if _, err := fmt.Scanf("%d", &c.Size); err != nil {
		return nil, err
	}

	var val int

	fmt.Print("Please specify the kind of matrix : \n\t1. Random\n\t2. Constant\n -> : ")
	if _, err := fmt.Scanf("%d", &val); err != nil {
		return nil, err
	} else if val == 2 {
		c.Type = Constant
	}

	return c, nil
}

//NewManualMatrixGenerator function will use settings given as parameter to create a MatrixGenerator
func NewManualMatrixGenerator(size int, t FillType) *MatrixGenerator {
	c := &MatrixGenerator{
		Size: size,
		Type: t,
	}
	return c
}

//Run function will allocate space for the matrix using config stored in MatrixGenerator.
func (m *MatrixGenerator) Run() {
	m.Matrix = make([][]int, m.Size)

	for i := 0; i < m.Size; i++ {
		m.Matrix[i] = make([]int, m.Size)

		for j := 0; j < m.Size; j++ {
			if m.Type == Random {
				m.Matrix[i][j] = rand.Intn(lapjv.MaxValue)
			} else {
				m.Matrix[i][j] = i * j % lapjv.MaxValue
			}
		}
	}
}

//Save function takes an io.Writer and saves the matrix in MatrixGenerator.Matrix to it.
func (m *MatrixGenerator) Save(out io.Writer) error {
	enc, err := json.Marshal(m.Matrix)
	if err != nil {
		return err
	}
	out.Write(enc)
	return nil
}
