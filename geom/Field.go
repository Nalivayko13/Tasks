package geom

import (
	"fmt"
)

type Field struct {
	n int
	f [][]int
}

func CreateField(n int) *Field {
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
	fmt.Println("--------------------")
}

func (f *Field) CleanField(){
	for i:=0;i<f.n;i++{
		for j:=0;j<f.n;j++{
			f.f[i][j]=0
		}
	}

}


func (f *Field) SetThreePoints(p1, p2, p3 Point) {
	if f.n<p1.x || f.n<p1.y || f.n<p2.x || f.n<p2.y || f.n<p3.x || f.n<p3.y{
		//TODO add sign check
		fmt.Println("going out of bounds!")
	}else {
		f.f[p1.x][p1.y]=1
		f.f[p2.x][p2.y]=1
		f.f[p3.x][p3.y]=1
	}
}
func ChangePoint(p1,p2,p3 Point,f *Field) {
	f.CleanField()
	f.SetThreePoints(p1,p2,p3)
	p:= CreateParallelogram(p1,p2,p3,f)
	fmt.Println("P-am center is", p.Center)
	fmt.Println("P-am points: ", p.P1,p.P2,p.P3,p.P4)
	o:= CreateCircle(p)
	fmt.Println("Radius of circle = ",o.radius)
}

