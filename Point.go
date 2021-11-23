package main

type Point struct {
	x,y int
	Color
}

func CreatePoint(x,y int, c Color) Point{
	var p Point
	p.x=x
	p.y=y
	p.Color=c
	return p
}
func (p *Point) SetPointColor(c Color){
	p.Color= c
}
