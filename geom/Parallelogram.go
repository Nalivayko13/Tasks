package geom

import (
	"fmt"
	"math"
)

type Parallelogram struct {
	P1, P2, P3, P4 Point
	Center         Point
	square         float64
}

func CreateParallelogram(p1,p2,p3 Point, f *Field) Parallelogram {
	var pam Parallelogram
	pam.P3 =p3
	pam.P1 =p1
	pam.P2 =p2
	f.SetThreePoints(p1,p2,p3)
	pam.square= CalcSquare(p1,p2,p3)
	pam.P4 = Create4Point(p1,p2,p3,f)
	a,b := true, true
	pam.Center = FindCenter(p1,p2,&a,&b)
	return pam
}

func Create4Point(a,b,c Point, f *Field) Point {
	var d Point
	inc:=true
	inc2:=true
	center:= FindCenter(a,b,&inc,&inc2)
	d.X =2*center.X - c.X
	d.Y =2*center.Y - c.Y

	if inc==false{
		d.X +=1
	}
	if inc2==false{
		d.Y +=1
	}
	f.F[d.X][d.Y]=1
	return d
}

func FindCenter(p1, p2 Point,Inc,Inc2 *bool) (Point){
	var center Point
	*Inc=true
	*Inc2=true
	var var1,var2 float64
	var1= float64((p1.X + p2.X) )/ 2
	var2= float64((p1.Y + p2.Y) )/ 2
	center.X = (p1.X + p2.X) / 2
	center.Y = (p1.Y + p2.Y) / 2
	if var1-float64(center.X)!=0 {
		*Inc=false
	}
	if var2-float64(center.Y)!=0{
		*Inc2=false
	}

	return center
}



func CalcSquare(p1,p2,p3 Point) float64 {
	S:=math.Abs(float64(p1.X*(p2.Y-p3.Y)-p1.Y*(p2.X-p3.X)+p2.X*p3.Y -p2.Y*p3.X))
	fmt.Println("square = ", S)
	return S
}