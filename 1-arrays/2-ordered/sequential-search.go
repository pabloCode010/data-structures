package main

import (
	"fmt"
)

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 14.
func sequential_search(values *[100]int, length, element int) (int, error) {
	var i int = 0
	for i < length && values[i] < element {
		i = i + 1
	}
	if i >= length || values[i] > element {
		return -i, fmt.Errorf("element %d not found", element)
	}
	return i, nil
}

// func main() {
// 	var values = [100]int{3, 5, 7, 12}
// 	var element = 7
// 	var length = 0

// 	pos, err := sequential_search(&values, length, element)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println(pos)
// 	} else {
// 		fmt.Printf("Element %d found at position %d\n", element, pos)
// 	}
// }
