package main

import "fmt"

func main() {
	var array [3][3][3]int

	value := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				array[i][j][k] = value
				value++
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Print(array[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
