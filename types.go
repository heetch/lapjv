package lapjv

//MatriceRow is the type we use to define an entry of our Matrice
type MatriceRow []int

//Matrice is the type we use to define our matrice all over the program.
type Matrice []MatriceRow

//FillType is the way we will fill our matrice in the generator.
type FillType int

const (
	//Random will fill our matrice with random values using math.rand lib and values from 0 to MaxValue.
	Random FillType = iota
	//Constant will fill our matrice with constant values with value = i * j for each entry
	Constant FillType = iota
)
