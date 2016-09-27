package lapjv

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

//matriceConfig will store info we will use to generate our matrice and store it.
//These info are set using CLI - The `matriceGeneratorInit` function.
type matriceConfig struct {
	Size int
	Type FillType
}

//matriceGeneratorInitInteractive function will prompt the user for each config part we want to fill
//It will then fill a matriceConfig struct and return it.
func matriceGeneratorInitInteractive() *matriceConfig {
	c := &matriceConfig{Type: Random}

	fmt.Print("Please enter the size of your matrice : ")
	if _, err := fmt.Scanf("%d", &c.Size); err != nil {
		panic(err)
	}

	var val int

	fmt.Print("Please specify the king of matrice : \n\t1. Random\n\t2. Constant\n -> : ")
	if _, err := fmt.Scanf("%d", &val); err != nil {
		panic(err)
	} else if val == 2 {
		c.Type = Constant
	}

	return c
}

//matriceGeneratorInitManual function will use settings given as parameter to store the config.
//It will fill a matriceConfig struct and return it.
func matriceGeneratorInitManual(size int, t FillType) *matriceConfig {
	c := &matriceConfig{
		Size: size,
		Type: t,
	}
	return c
}

//matriceGeneratorRun function is called in interactive and manual mode and will allocate space for the matrice using config.
//The function will then fill the matrice and return it.
func matriceGeneratorRun(c *matriceConfig) Matrice {
	m := make(Matrice, c.Size)

	for i := 0; i < c.Size; i++ {
		m[i] = make(MatriceRow, c.Size)

		for j := 0; j < c.Size; j++ {
			if c.Type == Random {
				m[i][j] = rand.Intn(MaxValue)
			} else {
				m[i][j] = i * j % MaxValue
			}
		}
	}
	return m
}

//matriceGeneratorSave function take a matrice - already filled - and a io.Writer as parameter and use the writer to save the matrice.
func matriceGeneratorSave(m Matrice, out io.Writer) {
	enc, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	out.Write(enc)
}

//The main difference between the two functions below is on the matriceGeneratorInit function that fill the config.

//MatriceGeneratorInteractive generate a matrice and save it in the out variable.
//This function use the matriceGeneratorInitInteractive function to fill the config struct in prompt mode.
func MatriceGeneratorInteractive(out io.Writer) {
	c := matriceGeneratorInitInteractive()

	m := matriceGeneratorRun(c)

	matriceGeneratorSave(m, out)
}

//MatriceGeneratorManual generate a matrice and save it in the out variable.
//This function use parameters given (size and fillType) to fill the config struct.
func MatriceGeneratorManual(out io.Writer, size int, fillType FillType) {
	c := matriceGeneratorInitManual(size, fillType)

	m := matriceGeneratorRun(c)

	matriceGeneratorSave(m, out)
}
