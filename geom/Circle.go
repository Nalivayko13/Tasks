package geom

import (
	"fmt"
	"math"
)

type Circle struct {
	Center Point
	Length float64
	Radius float64
	Square float64
	Color
}

func CreateCircle(p Parallelogram) Circle {
	var c Circle
	c.Center =p.Center
	c.Square =p.Square
	c.CalcRadius(c.Square)
	c.Color=Color((p.Center.Y+p.Center.X)/2)
	return c
}


func (c *Circle) CalcRadius(s float64){
	c.Radius =math.Sqrt(s/math.Pi)
	c.Length =2*(math.Pi)*c.Radius
	fmt.Println("Radius = ",c.Radius," length = ", c.Length)
}
