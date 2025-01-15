# Arrays
## Introduction

An **array** is a data structure that stores a set of elements in contiguous memory locations. It is one of the most basic and essential data structures in programming, providing fast access to elements using indexes.

Aquí está la actualización con el cambio solicitado: 

## Types of Arrays

1. **Ordered Arrays:**
   - Elements can be stored in any order, but the array can be implemented to ensure the elements are stored in a particular order (e.g., numerical or lexicographical).
   - Example:
     ```
     [1, 3, 5, 7, 9] // Ascending order
     ["apple", "banana", "cherry"] // Lexicographical order
     ```

2. **Unordered Arrays:**
   - Elements can be stored in any order, and there is no enforcement of a specific order.
   - Example:
     ```
     [7, 2, 9, 4, 1] // No defined order
     ["dog", "cat", "bird"]
     ```


## Representation of Multidimensional Arrays

When working with multidimensional arrays (like matrices), they can be represented in memory as one-dimensional arrays. There are two main ways to represent them:

### 1. **Row-Major Order**
   - The elements of each row of the matrix are stored consecutively in memory.
   - Function to calculate the position in a linear array:
     ```go
     func getPositionRMO(columns, i, j int) int {
     	return columns*i + j
     }
     ```
   
   **Example:**
   For a 2D matrix of size 2x3:
   ```
   [ [1, 2, 3],
     [4, 5, 6] ]
   ```
   Memory representation (row-major):
   ```
   [1, 2, 3, 4, 5, 6]
   ```

### 2. **Column-Major Order**
   - The elements of each column of the matrix are stored consecutively in memory.
   - Function to calculate the position in a linear array:
     ```go
     func getPositionCMO(rows, i, j int) int {
     	return rows*j + i
     }
     ```
   
   **Example:**
   For a 2D matrix of size 2x3:
   ```
   [ [1, 2, 3],
     [4, 5, 6] ]
   ```
   Memory representation (column-major):
   ```
   [1, 4, 2, 5, 3, 6]
   ```