package main

import "fmt"

type Field struct {
	n int
	f [][]int
}

func CreateField(n int) *Field{
	var matrix Field
	matrix.n=n
	matrix.f = make([][]int, n)
	for i := range matrix.f {
		matrix.f[i] = make([]int, n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			matrix.f[i][j]=0
		}
	}

	return &matrix
}

func (field Field) PrintField(){
	for i:=0;i<field.n;i++{
		for j:=0;j<field.n;j++{
			fmt.Print(field.f[i][j]," ")
		}
		fmt.Println("")
	}
}


func (f *Field) SetThreePoints(p1, p2, p3 Point) {

}
