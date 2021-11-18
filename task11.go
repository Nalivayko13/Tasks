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

func ReverselineParentheses(s string) string{
	var Reverse = func(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	var (
		start, stop int
		out string
	)

	for i:=0;i<len(s);i++{
		for j:=0;j<len(s);j++{
			if s[j]=='(' {
				start=j
				break
			}
		}
		for j:=0;j<len(s);j++{
			if s[j]==')' {
				stop=j
				break
			}
		}
		out=s[:start]+Reverse(s[start+1:stop])+s[stop+1:]
		//fmt.Println(start)
		//fmt.Println(stop)

	}
	return out
}

func main() {
	//fmt.Println(CheckPalindrom("aabaa"))
	//fmt.Println(CheckPalindrom("aabba"))
	//fmt.Println(MatrixElementsSum(3))
	fmt.Println(ReverselineParentheses("foo(bar)baz"))
	fmt.Println(ReverselineParentheses("foo(bar)baz(blim)"))
	fmt.Println(ReverselineParentheses("foo(bar(baz))blim"))
	fmt.Println(ReverselineParentheses("foo(bar(baz))blim"))
}
