package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 15.
func insertElement(values *[100]int, length *int, newElement int) error {
	if *length < 100 {
		pos, err := sequential_search(values, *length, newElement)
		if err == nil && pos >= 0 {
			return fmt.Errorf("Element %d already exists", newElement)
		} else {
			*length = *length + 1
			pos = pos * -1
			for i := *length - 1; i >= pos+1; i-- {
				values[i] = values[i-1]
			}
			values[pos] = newElement
		}
	} else {
		return fmt.Errorf("Array is full")
	}
	return nil
}

// func main() {
// 	var values [100]int
// 	var length = 0
// 	var newElement = 1
// 	fmt.Printf("%v\n", values[:length])
// 	insertElement(&values, &length, newElement)

// 	fmt.Printf("%v\n", values[:length])
// }

// to run use:
// $ go run sequential-search.go insert-element.go
