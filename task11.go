package main

import (
	"fmt"
	"math/rand"
)

func CheckPalindrom(s string) bool{
	for i:=0;i<len(s)/2;i++{
		if s[i]!=s[len(s)-i-1]{
			return false
		}
	}
	return true
}

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

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(CheckPalindrom("aabaa"))
	fmt.Println(CheckPalindrom("aabba"))

	fmt.Println(MatrixElementsSum(3))
}
