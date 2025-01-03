package main

import (
	"errors"
)

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 11.
func insertElement(values *[100]int, length *int, newElement int) error {
	if *length < 100 {
		values[*length] = newElement
		*length = *length + 1
		return nil
	} else {
		return errors.New("array is full")
	}
}

// func main() {
// 	var values = [100]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	var length = 10
// 	insertElement(&values, &length, 11)
// 	insertElement(&values, &length, 12)
// 	fmt.Println(values[:length])
// }
