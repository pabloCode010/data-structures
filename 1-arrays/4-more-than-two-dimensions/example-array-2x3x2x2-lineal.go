package main

// import "fmt"

// func getPositionRMO(depth, rows, columns, i, j, k, l int) int {
// 	return i*depth*rows*columns + j*rows*columns + k*columns + l
// 	//POS = k1 * (l2 * l3 * l4) + k2 * (l3 * l4) + k3 * l4 + k4
// }

// func getPositionCMO(layers, depth, rows, i, j, k, l int) int {
// 	return l*layers*depth*rows + k*layers*depth + j*layers + i
// 	//POS = k4 * (l1 * l2 * l3) + k3 * (l1 * l2) + k2 * l1 + k1
// }

// func main() {
// 	var layers = 2  // l1
// 	var depth = 3   // l2
// 	var rows = 2    // l3
// 	var columns = 2 // l4

// 	var arrayRMO [24]int // 2 * 3 * 2 * 2 = 24
// 	var arrayCMO [24]int

// 	value := 0
// 	for i := 0; i < layers; i++ {
// 		for j := 0; j < depth; j++ {
// 			for k := 0; k < rows; k++ {
// 				for l := 0; l < columns; l++ {
// 					posRMO := getPositionRMO(depth, rows, columns, i, j, k, l)
// 					posCMO := getPositionCMO(layers, depth, rows, i, j, k, l)
// 					arrayRMO[posRMO] = value
// 					arrayCMO[posCMO] = value
// 					value++
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < layers; i++ {
// 		for j := 0; j < depth; j++ {
// 			for k := 0; k < rows; k++ {
// 				for l := 0; l < columns; l++ {
// 					posRMO := getPositionRMO(depth, rows, columns, i, j, k, l)
// 					posCMO := getPositionCMO(layers, depth, rows, i, j, k, l)

// 					fmt.Printf("index [%d][%d][%d][%d] -> RMO: %d, CMO: %d\n",
// 						i, j, k, l, arrayRMO[posRMO], arrayCMO[posCMO])
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("Row Major Order   :", arrayRMO)
// 	fmt.Println("Column Major Order:", arrayCMO)
// }
