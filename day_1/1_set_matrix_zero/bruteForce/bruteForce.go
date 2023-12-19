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

func makeRowAndColumnMinusOne(matrix [][]int, r, c int) [][]int {
	// make row zero
	for i := 0; i < n; i++ {
		matrix[r][i] = -1
	}

	//make column zero
	for i := 0; i < n; i++ {
		matrix[i][c] = -1
	}

	return matrix
}

func makeMinusOneToZero(matrix [][]int) [][]int {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

func main() {
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
				matrix = makeRowAndColumnMinusOne(matrix, i, j)
			}
		}
	}

	matrix = makeMinusOneToZero(matrix)

	// print result
	printMatrix(matrix)

}
