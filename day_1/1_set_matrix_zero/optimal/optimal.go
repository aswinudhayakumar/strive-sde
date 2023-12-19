package optimal

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
	col1 := 1
	matrix := make([][]int, n, n)
	matrix = [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	// col[n] = matrix[0][...]
	// row[n] = matrix[...][0]

	// print input
	printMatrix(matrix)

	// process
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {

				matrix[i][0] = 0

				if j != 0 {
					matrix[0][j] = 0
				} else {
					col1 = 0
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != 0 {
				if matrix[i][0] == 0 || matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}

		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if col1 == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

	printMatrix(matrix)
}
