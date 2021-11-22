package main

type Point struct {
	x,y int
	Color
}
func (p *Point) SetPointColor(c Color){
	p.Color= c
}
