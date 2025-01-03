package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 13.
func updateElement(values *[100]int, length, element, newElement int) error {
	var i int = 0
	for i < length && element != values[i] {
		i = i + 1
	}
	if i >= length {
		return fmt.Errorf("element %d not found", element)
	} else {
		values[i] = newElement
		return nil
	}
}

// func main() {
// 	var (
// 		values     = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 		length     = 10
// 		element    = 5
// 		newElement = 99
// 	)
// 	err := updateElement(&values, length, element, newElement)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(values[:length])
// 	}
// }
