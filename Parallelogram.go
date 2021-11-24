package main

import (
	"fmt"
	"math"
)

type Parallelogram struct {
	p1, p2, p3, p4 Point
}
func Create4Point(a,b,c Point) Point{
	inc:=true
	inc2:=true
	//TODO out of bounds check
	var d Point
	r1:=math.Sqrt(math.Pow(float64((a.x - b.x)),2)+math.Pow(float64((a.y-b.y)),2))
	fmt.Println(r1)
	r2:=math.Sqrt(math.Pow(float64((b.x - c.x)),2)+math.Pow(float64((b.y-c.y)),2))
	fmt.Println(r2)
	r3:=math.Sqrt(math.Pow(float64((a.x - c.x)),2)+math.Pow(float64((a.y-c.y)),2))
	fmt.Println(r3)
	max:=math.Max(r1, math.Max(r2,r3))

	if r1!=r2 && r2!=r3 && r1!=r3 {
		switch max {
		case r1:
			center := FindCenter(a, b, &inc, &inc2)
			d.x = 2*center.x - c.x
			d.y = 2*center.y - c.y
		case r2:
			center := FindCenter(b, c, &inc,&inc2)
			d.x = 2*center.x - a.x
			d.y = 2*center.y - a.y
		case r3:
			center := FindCenter(a, c, &inc,&inc2)
			d.x = 2*center.x - b.x
			d.y = 2*center.y - b.y
		default:
			fmt.Println("oh nooo")
		}
	}else {fmt.Println("Габэлла")
		if r1==r2{
			if center := FindCenter(a, b, &inc,&inc2);center.x<0 && center.y<0 {
				d.x = 2*center.x - c.x
				d.y = 2*center.y - c.y
			}else {
				center := FindCenter(b, c, &inc,&inc2)
				d.x = 2*center.x - a.x
				d.y = 2*center.y - a.y
			}
		}
		if r2==r3{
			if center := FindCenter(b, c, &inc,&inc2); center.x<0 && center.y<0{

				d.x = 2*center.x - a.x
				d.y = 2*center.y - a.y
			}else {
				center := FindCenter(a, c,&inc,&inc2)
				d.x = 2*center.x - b.x
				d.y = 2*center.y - b.y
			}
		}
		if r1==r3{

			if center := FindCenter(a, c,&inc,&inc2); center.x<0 && center.y<0{
				d.x = 2*center.x - b.x
				d.y = 2*center.y - b.y
			}else {
				center := FindCenter(a, b,&inc,&inc2)
				d.x = 2*center.x - c.x
				d.y = 2*center.y - c.y
			}
		}
	}

	if inc==false{
		d.x+=1
	}
	if inc2==false{
		d.y+=1
	}

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

func (pam Parallelogram) CalcSquare() float64{

	return 0
}