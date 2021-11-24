package main

import (
	"fmt"
	"math"
)

type Parallelogram struct {
	p1, p2, p3, p4 Point
	center Point
	square float64
}

func CreateParallelogram(p1,p2,p3 Point, f *Field) Parallelogram{
	var pam Parallelogram
	pam.p3=p3
	pam.p1=p1
	pam.p2=p2
	f.SetThreePoints(p1,p2,p3)
	pam.square=CalcSquare(p1,p2,p3)
	pam.p4=Create4Point(p1,p2,p3,f)
	a,b := true, true
	pam.center=FindCenter(p1,p2,&a,&b)
	return pam
}

func Create4Point(a,b,c Point, f *Field) Point{
	var d Point
	inc:=true
	inc2:=true
	center:=FindCenter(a,b,&inc,&inc2)
	d.x=2*center.x - c.x
	d.y=2*center.y - c.y

	if inc==false{
		d.x+=1
	}
	if inc2==false{
		d.y+=1
	}
	f.f[d.x][d.y]=1
	return d
}

func FindCenter(p1, p2 Point,Inc,Inc2 *bool) (Point){
	var center Point
	*Inc=true
	*Inc2=true
	var var1,var2 float64
	var1= float64((p1.x + p2.x) )/ 2
	var2= float64((p1.y + p2.y) )/ 2
	center.x = (p1.x + p2.x) / 2
	center.y = (p1.y + p2.y) / 2
	if var1-float64(center.x)!=0 {
		*Inc=false
	}
	if var2-float64(center.y)!=0{
		*Inc2=false
	}

	return center
}



func CalcSquare(p1,p2,p3 Point) float64 {
	S:=math.Abs(float64(p1.x*(p2.y-p3.y)-p1.y*(p2.x-p3.x)+p2.x*p3.y-p2.y*p3.x))
	fmt.Println("square = ", S)
	return S
}