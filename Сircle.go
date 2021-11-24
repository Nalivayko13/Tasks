package main

import (
	"fmt"
	"math"
)

type Circle struct {
	center Point
	length float64
	radius float64
	square float64
	Color
}

func CreateCircle(p Parallelogram) Circle{
	var c Circle
	c.center=p.center
	c.square=p.square
	c.CalcRadius(c.square)
	//c.Color=Color()
	return c
}

func (circ *Circle) SetCircleColor(c Color){
	circ.Color= c
}

func (c *Circle) CalcRadius(s float64){
	c.radius=math.Sqrt(s/math.Pi)
	c.length=2*(math.Pi)*c.radius
	fmt.Println("Radius = ",c.radius," length = ", c.length)
}
