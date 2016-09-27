package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/heetch/lapjv"
	"github.com/spf13/cobra"
)

const (
	//resourcesStorage will be used in our util functions that read / write to files.
	resourcesStorage = "resources/"
)

var (
	filename    string
	size        int
	Constant    string
	interactive bool
)

//RootCmd is the main command displayed by Cobra with no argument
var RootCmd = &cobra.Command{
	Use:   "LAPJV Algorithm - Golang implementation",
	Short: "An implementation of the LAPJV Algorithm working in Golang.",
	Long:  "The Linear Assignment Problem Solver by Jonker-Volgenant Algorithm - with benchmark and test.",
}

var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generate a JSON file that describe the matrice with given parameters.",
	Long:  "Use this command and generate a JSON file that describe the matrice you want to resolve - will be saved in 'resources' folder.",
	Run:   runGenerator,
}

//runGenerator function will be called in order to generate a matrice and save it in a file.
//This function create the file and run a function between MatriceGeneratorInteractive and MatriceGeneratorManual following the CLI flags.
func runGenerator(cmd *cobra.Command, args []string) {

	if filename == "" {
		filename = "example.out"
	}
	f, err := os.Create(resourcesStorage + filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if interactive == true {
		lapjv.MatriceGeneratorInteractive(f)
	} else {
		t := lapjv.Random
		if Constant == "constant" {
			t = lapjv.Constant
		}

		lapjv.MatriceGeneratorManual(f, size, t)
	}
}

var solverCmd = &cobra.Command{
	Use:   "solver",
	Short: "Solve a matrice described in the JSON file given as parameter",
	Long:  "Use this command to solve a matrice you described in the JSON file before. Response will be printed in stdout",
	Run:   runSolver,
}

//runSolver function will be called in order to solve the matrice using a file or a generated matrice using the generator.
//This function open the file and read the content. Once this step done, it will call the MatriceSolver itself.
func runSolver(cmd *cobra.Command, args []string) {

	buf := new(bytes.Buffer)

	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		buf.ReadFrom(f)
	} else if interactive == true {
		lapjv.MatriceGeneratorInteractive(buf)
	} else {
		t := lapjv.Random
		if Constant == "constant" {
			t = lapjv.Constant
		}

		lapjv.MatriceGeneratorManual(buf, size, t)
	}

	var m lapjv.Matrice
	if err := json.Unmarshal(buf.Bytes(), &m); err != nil {
		panic(err)
	}
	lapjv.MatriceSolver(m)
}

func init() {
	//Set flags to the program CLI Commands
	RootCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "file in which the matrice will be stored")
	RootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Set the value to true in order to run the generator in interactive mode")
	RootCmd.PersistentFlags().StringVarP(&Constant, "type", "t", "worst", "Set the value to true in order to fill the matrice with Constant case values (between worst and constant)")
	RootCmd.PersistentFlags().IntVarP(&size, "size", "s", 10, "size of the matrice.")

	//Add commands to the program CLI
	RootCmd.AddCommand(generatorCmd, solverCmd)
}

//Function used by Cobra to execute the command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(-1)
	}
}
