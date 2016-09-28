package lapjv

const (
	//MaxValue should not be changed.
	//This Value permit to establish a Max Value we can give to an enty of the matrix.
	MaxValue = 100000
)

//MatrixSolver function take a Matrix - already filled -  as parameter and declare useful variables for the Lapjv algo itself.
//After this first step, it call the Lapjv algorithm and save the result.
func MatrixSolver(m [][]int) int {

	rowsol := make([]int, len(m[0]))
	colsol := make([]int, len(m[0]))
	u := make([]int, len(m[0]))
	v := make([]int, len(m[0]))

	cost := Lapjv(len(m[0]), m, rowsol, colsol, u, v)
	return cost
}

// Lapjv is a naive port of the Jonker Volgenant Algorithm from C++ to Go
func Lapjv(dim int, assigncost [][]int, rowsol, colsol, u, v []int) int {
	var unassignedfound bool
	var i, imin, numfree, prvnumfree, i0, freerow int
	var j, j1, j2, endofpath, last, low, up int
	var min, h, umin, usubmin, v2 int

	free := make([]int, dim)
	collist := make([]int, dim)
	matches := make([]int, dim)
	pred := make([]int, dim)
	d := make([]int, dim)

	// skipping L53-54
	for j := dim - 1; j >= 0; j-- {
		min = assigncost[0][j]
		imin = 0
		for i := 1; i < dim; i++ {
			if assigncost[i][j] < min {
				min = assigncost[i][j]
				imin = i
			}
		}

		v[j] = min
		matches[imin]++
		if matches[imin] == 1 {
			rowsol[imin] = j
			colsol[j] = imin
		} else {
			colsol[j] = -1
		}
	}

	for i := 0; i < dim; i++ {
		if matches[i] == 0 {
			free[numfree] = i
			numfree++
		} else if matches[i] == 1 {
			j1 = rowsol[i]
			min = MaxValue
			for j := 0; j < dim; j++ {
				if j != j1 && assigncost[i][j]-v[j] < min {
					min = assigncost[i][j] - v[j]
				}
			}
			v[j1] -= min
		}
	}

	for loopcmt := 0; loopcmt < 2; loopcmt++ {
		k := 0
		prvnumfree = numfree
		numfree = 0
		for k < prvnumfree {
			i = free[k]
			k++
			umin = assigncost[i][0] - v[0]
			j1 = 0
			usubmin = MaxValue

			for j := 1; j < dim; j++ {
				h = assigncost[i][j] - v[j]

				if h < usubmin {
					if h >= umin {
						usubmin = h
						j2 = j
					} else {
						usubmin = umin
						umin = h
						j2 = j1
						j1 = j
					}
				}
			}

			i0 = colsol[j1]
			if umin < usubmin {
				v[j1] = v[j1] - (usubmin - umin)
			} else if i0 >= 0 {
				j1 = j2
				i0 = colsol[j2]
			}

			rowsol[i] = j1
			colsol[j1] = i
			if i0 >= 0 {
				if umin < usubmin {
					k--
					free[k] = i0
				} else {
					free[numfree] = i0
					numfree++
				}
			}
		}
	}

	for f := 0; f < numfree; f++ {
		freerow = free[f]
		for j := 0; j < dim; j++ {
			d[j] = assigncost[freerow][j] - v[j]
			pred[j] = freerow
			collist[j] = j
		}

		low = 0
		up = 0
		unassignedfound = false

		for !unassignedfound {
			if up == low {
				last = low - 1
				min = d[collist[up]]
				up++

				for k := up; k < dim; k++ {
					j = collist[k]
					h = d[j]
					if h <= min {
						if h < min {
							up = low
							min = h
						}
						collist[k] = collist[up]
						collist[up] = j
						up++
					}
				}

				for k := low; k < up; k++ {
					if colsol[collist[k]] < 0 {
						endofpath = collist[k]
						unassignedfound = true
						break
					}
				}
			}

			if !unassignedfound {
				j1 = collist[low]
				low++
				i = colsol[j1]
				h = assigncost[i][j1] - v[j1] - min

				for k := up; k < dim; k++ {
					j = collist[k]
					v2 = assigncost[i][j] - v[j] - h

					if v2 < d[j] {
						pred[j] = i

						if v2 == min {
							if colsol[j] < 0 {
								endofpath = j
								unassignedfound = true
								break
							} else {
								collist[k] = collist[up]
								collist[up] = j
								up++
							}
						}

						d[j] = v2
					}
				}
			}
		}

		for k := 0; k <= last; k++ {
			j1 = collist[k]
			v[j1] += d[j1] - min
		}

		i = freerow + 1
		for i != freerow {
			i = pred[endofpath]
			colsol[endofpath] = i
			j1 = endofpath
			endofpath = rowsol[i]
			rowsol[i] = j1
		}
	}

	lapcost := 0
	for i := 0; i < dim; i++ {
		j = rowsol[i]
		u[i] = assigncost[i][j] - v[j]
		lapcost += assigncost[i][j]
	}

	return lapcost
}
