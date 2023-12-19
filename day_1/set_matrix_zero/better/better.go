package main

import "fmt"

var n = 3

func printMatrix(matrix [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	col := make([]int, n)
	row := make([]int, n)

	matrix := make([][]int, n, n)
	matrix = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	// print input
	printMatrix(matrix)

	// process
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				col[i], row[j] = 1, 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if col[i] == 1 || row[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}

	// print result
	printMatrix(matrix)
}
