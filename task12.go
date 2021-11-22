package main

import (
	"fmt"
	"math/rand"
)

func MatrixElementsSum(n int) int{
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			matrix[i][j]=rand.Intn(2)
		}
	}
	var PrintMatrix = func(matrix [][]int){
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(matrix[i][j], " ")
			}
			fmt.Println(" ")
		}
	}
	PrintMatrix(matrix)
	sum:=0
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if matrix[j][i]==0 {
				break
			}else {
				sum += matrix[j][i]
			}
		}
	}

	return sum
}


