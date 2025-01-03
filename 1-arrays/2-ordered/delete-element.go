package main

import "fmt"

// This implementation is based on the algorithm described in
// "Estructuras de datos Tercera edición" by [Osvaldo Cairó Silvia Guardati], Page 16.
func deleteElement(values *[100]int, length *int, element int) error {
	if *length > 0 {
		pos, err := sequential_search(values, *length, element)
		if err != nil {
			return err
		} else {
			*length = *length - 1
			for i := pos; i < *length; i++ {
				values[i] = values[i+1]
			}
		}
	} else {
		return fmt.Errorf("Array is empty")
	}
	return nil
}

// func main() {
// 	var values = [100]int{2, 4, 6, 8, 10}
// 	var length = 5
// 	var element = 6

// 	fmt.Printf("%v\n", values[:length])
// 	err := deleteElement(&values, &length, element)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%v\n", values[:length])
// 		fmt.Printf("All Spaces: %v\n", values)
// 		// Length: 4 |
// 		// [2 4 8 10 | 10]
// 		// The last element is duplicated
// 	}
// }

// // to run use:
// // $ go run sequential-search.go delete-element.go
