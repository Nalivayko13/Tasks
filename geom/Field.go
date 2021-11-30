package geom

import (
	"fmt"
)

type Field struct {
	N int
	F [][]int
}

func CreateField(n int) Field {
	var matrix Field
	matrix.N =n
	matrix.F = make([][]int, n)
	for i := range matrix.F {
		matrix.F[i] = make([]int, n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			matrix.F[i][j]=0
		}
	}

	return matrix
}

func (field Field) PrintField(){
	for i:=0;i<field.N;i++{
		for j:=0;j<field.N;j++{
			fmt.Print(field.F[i][j]," ")
		}
		fmt.Println("")
	}
	fmt.Println("--------------------")
}

func (f *Field) CleanField(){
	for i:=0;i<f.N;i++{
		for j:=0;j<f.N;j++{
			f.F[i][j]=0
		}
	}

}


func (f *Field) SetThreePoints(p1, p2, p3 Point) {
	if f.N <p1.X || f.N <p1.Y || f.N <p2.X || f.N <p2.Y || f.N <p3.X || f.N <p3.Y {
		//TODO add sign check
		fmt.Println("going out of bounds!")
	}else {
		f.F[p1.X][p1.Y]=1
		f.F[p2.X][p2.Y]=1
		f.F[p3.X][p3.Y]=1
	}
}
func ChangePoint(p1,p2,p3 Point,f *Field) {
	f.CleanField()
	f.SetThreePoints(p1,p2,p3)
	p:= CreateParallelogram(p1,p2,p3,f)
	fmt.Println("P-am center is", p.Center)
	fmt.Println("P-am points: ", p.P1,p.P2,p.P3,p.P4)
	o:= CreateCircle(p)
	fmt.Println("Radius of circle = ",o.Radius)
}

