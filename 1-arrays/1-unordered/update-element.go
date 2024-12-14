package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 13.
func updateElement(values *[100]int, length, element, newElement int) {
	var i int = 0
	for i < length && element != values[i] {
		i++
	}
	if i >= length {
		fmt.Printf("Element %d not found\n", element)
	} else {
		values[i] = newElement
	}
}

// func main() {
// 	var (
// 		values     = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 		length     = 10
// 		element    = 5
// 		newElement = 99
// 	)
// 	updateElement(&values, length, element, newElement)
// 	fmt.Println(values)
// }
