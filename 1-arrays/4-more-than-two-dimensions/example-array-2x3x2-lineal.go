package main

import "fmt"

func getPositionRMO(rows, columns, i, j, k int) int {
	return i*rows*columns + j*columns + k
}

func getPositionCMO(layers, rows, i, j, k int) int {
	return k*layers*rows + j*layers + i
}

func main() {
	var arrayRMO [12]int
	var arrayCMO [12]int

	var layers = 2
	var rows = 3
	var columns = 2

	value := 0
	for i := 0; i < layers; i++ {
		for j := 0; j < rows; j++ {
			for k := 0; k < columns; k++ {
				posRMO := getPositionRMO(rows, columns, i, j, k)
				posCMO := getPositionCMO(layers, rows, i, j, k)
				arrayRMO[posRMO] = value
				arrayCMO[posCMO] = value
				value++
			}
		}
	}

	// get the same position in both arrays
	for i := 0; i < layers; i++ {
		for j := 0; j < rows; j++ {
			for k := 0; k < columns; k++ {
				posRMO := getPositionRMO(rows, columns, i, j, k)
				posCMO := getPositionCMO(layers, rows, i, j, k)

				fmt.Println("index [", i, "][", j, "][", k, "]", "RMO:", arrayRMO[posRMO], "CMO:", arrayCMO[posCMO])
			}
		}
	}

	fmt.Println("Row Major Order   :", arrayRMO)
	fmt.Println("Column Major Order:", arrayCMO)
}
