package main

import (
	"fmt"
	"math"
)
type Parallelogram struct {
	p1, p2, p3, p4 Point
}
func Create4Point(a,b,c Point) Point{
	var d Point
	r1:=math.Sqrt(math.Pow(float64((a.x - b.x)),2)+math.Pow(float64((a.y-b.y)),2))
	r2:=math.Sqrt(math.Pow(float64((b.x - c.x)),2)+math.Pow(float64((b.y-c.y)),2))
	r3:=math.Sqrt(math.Pow(float64((a.x - c.x)),2)+math.Pow(float64((a.y-c.y)),2))
	max:=math.Max(r1, math.Max(r2,r3))
	switch max {
	case r1:
		center:=FindCenter(a,b)
		d.x=2*center.x-c.x
		d.y=2*center.y-c.y
	case r2:
		center:=FindCenter(b,c)
		d.x=2*center.x-a.x
		d.y=2*center.y-a.y
	case r3:
		center:=FindCenter(a,c)
		d.x=2*center.x-b.x
		d.y=2*center.y-b.y
	default:
		fmt.Println("oh nooo")
	}
	return d
}

func FindCenter(p1, p2 Point) Point{
	var center Point
	center.x=(p1.x+p2.x)/2
	center.y=(p1.y+p2.y)/2
	return center
}

func (pam Parallelogram) CalcSquare() float64{

	return 0
}