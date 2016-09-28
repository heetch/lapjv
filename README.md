**LAPJV - Go Implementation**
======================

This repository is the Golang implementation of the **LAPJV** algorithm and includes utilitaries to facilitate the first usage and tests of the library.

----------
**Overview**
-------------

The repository contains the library that solve a matrix using the LAPJV algorithm.
It also contains a wrapper divided in two : 

 - A generator that generates a matrix and save it on a file.
 - A solver using a file with a matrix into it and.

These two commands are available over the CLI. The solver can also be used directly as explained below.

**Installation**
-------------

The library can be installed using the `go get` command : 

` > go get github.com/heetch/lapjv/...`

### Library

Just write, into your project : 

`import "github.com/heetch/lapjv`

### Tool

You now can use the `lapjv` command and see the usage.
*(If nothing appears check that the `$GOPATH/bin` folder is in your PATH env variable.*

**Getting Started**
-------------

### Using the Library

```
package main

import "github.com/heetch/lapjv"

func main() {
     m := make([][]int, 10)

     for i := 0; i < 10; i++ {
     	 m[i] = make([]int, 10)
	 }

		//You should fill your matrix here with cost values

		res := lapjv.SolveMatrix(m)

		//Here you now can use fields of the res variables to check the result.
		//Fields are :
		// - cost : the cost of the resolution
		// - rowsol : the solution, based on rows.
		// - colsol : the solution, based on cols.
}
```

### Using the Tool

Usage :
[Picture]

##### Generator
The generator can be used, in the simplest way, with : 
``` lapjv generator ``` . This command will generates a file named `example.out` with the JSON format of a 10*10 randomly filled matrix.

###### Interactive mode

You can use the interactive mode to specify options of the matrix you want to generate : `lapjv generator -i`. This mode also generates an `example.out` file in the current directory.

###### Manual mode

You can use the manual mode to specify options of the matrix you want to generate : `lapjv generator -s size -t constant`. This mode generates an `example.out` file in the current directory using only options given as parameter.

###### Specifying an ouput file

You can specify the file in which you want to write the matrix using the -f option. This option is available in both manual and interactive modes. You can do it with : `lapjv generator -f filename` This one option can be combined with ones above.

##### Solver

The solver can be used using : `lapjv solver` command.

You can launch it without any option. In this case , the solver with try to open the `example.out` file in the current directory in present, or just quit if the file is not present.

###### Using with the generator

You can generate a matrix and solve it in the same time using options of the generator.

Example : `lapjv solver -i` Will prompt you for options and solve the matrix just after.


**Error ?**
-------------
 If you discover any error, please create an issue describing it.

**License**
-------------
 The library is released under the MIT license. See LICENSE file.
