package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 10.
func sequential_search(values *[100]int, length, element int) (int, error) {
	var i int = 0
	for i < length && element != values[i] {
		i = i + 1
	}
	if i >= length {
		return -1, fmt.Errorf("element %d not found", element)
	}
	return i, nil
}

// func main() {
// 	var values = [100]int{7, 22, 8, 2, 0, 4, 17}
// 	var length = 10
// 	position, err := sequential_search(&values, length, 4)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("Element found at position %d\n", position)
// 	}
// 	position, err = sequential_search(&values, length, 99)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("Element found at position %d\n", position)
// 	}
// }
