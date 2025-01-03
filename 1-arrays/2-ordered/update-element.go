package main

import "fmt"

// This is my implementation of the algorithm to update an element in the array, which is left as task in the book
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 16
// You can read the instructions on that page.
func updateElement(values *[100]int, length *int, element, newElement int) error {
	pos, err := sequential_search(values, *length, element)
	if err != nil {
		return err
	}
	if (pos > 0 && values[pos-1] >= newElement) || (pos < *length-1 && values[pos+1] <= newElement) {
		if err := deleteElement(values, length, element); err != nil {
			return err
		}
		if err := insertElement(values, length, newElement); err != nil {
			return err
		}
	} else {
		values[pos] = newElement
	}
	return nil
}

func main() {
	var values = [100]int{2, 4, 8, 12}
	var length = 4
	var element = 4

	err := updateElement(&values, &length, element, 18)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(values[:length])
	}
}
