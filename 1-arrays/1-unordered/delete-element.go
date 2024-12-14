package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 12.
func deleteElement(values *[100]int, length *int, element int) {
	var i int = 0
	for i < *length && element != values[i] {
		i++
	}
	if i >= *length {
		fmt.Printf("Element %d not found\n", element)
	} else {
		for j := i; j < *length-1; j++ {
			values[j] = values[j+1]
		}
		*length = *length - 1
	}
}

// func main() {
// 	var values = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	var length = 10
// 	deleteElement(&values, &length, 5)
// 	fmt.Printf("Length: %d\n", length)
// 	fmt.Println(values)

// 	// the element 5 is deleted from the array and the length is decremented by 1, the last element is duplicated
// 	// Length: 9 _________|
// 	// [1 2 3 4 6 7 8 9 10|10]
// }
