package main

import "fmt"

func getPositionRMO(columns, i, j int) int {
	return columns*i + j
}

func getPositionCMO(rows, i, j int) int {
	return rows*j + i
}

func main() {
	var arrayRMO [9]int
	var arrayCMO [9]int
	var columns = 3
	var rows = 3
	var count int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			posRMO := getPositionRMO(columns, i, j)
			posCMO := getPositionCMO(rows, i, j)
			arrayRMO[posRMO] = count
			arrayCMO[posCMO] = count
			count++
		}
	}

	// array row major order
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arrayRMO[getPositionRMO(columns, i, j)], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	// array column major order
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arrayCMO[getPositionCMO(rows, i, j)], " ")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Row Major Order   :", arrayRMO)
	fmt.Println("Column Major Order:", arrayCMO)
}

// This program demonstrates how to represent a 2D matrix as a 1D array
// using two storage orders: Row Major Order (RMO) and Column Major Order (CMO).
//
// - Row Major Order (RMO): Stores matrix elements row by row in a 1D array.
// - Column Major Order (CMO): Stores matrix elements column by column in a 1D array.
//
// The program uses two helper functions to calculate the index for RMO and CMO
// based on the matrix dimensions and positions.
// It fills two arrays (arrayRMO and arrayCMO) with sequential values,
// then prints the matrix representation for both orders and their 1D array contents.
