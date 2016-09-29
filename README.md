**LAPJV - Go Implementation**
======================

This repository is the Golang implementation of the **LAPJV** algorithm and includes utilitaries to facilitate the first usage and tests of the library.

----------
**Overview**
-------------

The repository contains the library that solves a matrix using the LAPJV algorithm.
It also contains a wrapper divided in two : 

 - A generator that generates a matrix and saves it on a file.
 - A solver that uses a file containing a matrix.

These two commands are available over the CLI. The solver can also be used directly as explained below.

**Installation**
-------------

The library can be installed using the `go get` command : 

` > go get github.com/heetch/lapjv/...`

### Library

Include Lapjv in your project : 

`
import "github.com/heetch/lapjv"`

### Tool

You now can use the `lapjv` command and see the usage.
*(If nothing appears check that the `$GOPATH/bin` folder is in your PATH env variable.*

**Getting Started**
-------------

### Using the Library

```go
package main

import (
        "fmt"
        "math/rand"

        "github.com/heetch/lapjv"
)

func main() {

     // Create your matrix here and fill it with values.
     m := make([][]int, 10)

     for i := 0; i < 10; i++ {
     	 m[i] = make([]int, 10)
         for j := 0; j < 10; j++ {
        		//You could fill your matrix here with cost values

                m[i][j] = rand.Intn(10000)
         }
	 }

		res := lapjv.MatrixSolver(m)

		//Here you now can use fields of the res variable to check the result.
		//Fields are :
		// - Cost : the cost of the resolution
		// - Rowsol : the solution, based on rows.
		// - Colsol : the solution, based on cols.

		fmt.Println(res.Cost)
}
```

### Using the Tool

Usage :
[Picture]

##### Generator
The generator can be used, in the simplest way, with : 

``` lapjv generator ```

This command will generate a file named `example.json` with the JSON format of a 10*10 randomly filled matrix.

###### Interactive mode

You can use the interactive mode to specify options of the matrix you want to generate : 

```lapjv generator -i```

This mode also generates an `example.json` file in the current directory.

###### Manual mode

You can use the manual mode to specify options of the matrix you want to generate : 

```lapjv generator -s size -t constant```

This mode generates an `example.json` file in the current directory using only options given as parameter.

###### Specifying an output file

You can specify the file in which you want to write the matrix using the -f option. This option is available in both manual and interactive modes. You can do it with :

```lapjv generator -f filename```

This one option can be combined with ones above.

##### Solver

The solver can be used using :

```lapjv solver```

You can launch it without any option. In this case , the solver will read from the stdin, waiting for a JSON-formated matrix.

###### Using with the generator

You can generate a matrix and solve it in the same time using options of the generator.

Example : 

```lapjv solver -i``` 

This command will prompt you for options and solve the matrix just after.


**Error ?**
-------------
 If you discover any error, please create an issue describing it.

**License**
-------------
 The library is released under the MIT license. See LICENSE file.
