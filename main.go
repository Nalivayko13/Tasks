package main

import "fmt"

const (
	WHITE= iota
	BLACK
	PINK
	RED
	GREEN
	BROWN
	ORANGE
	YELLOW
)

type Color byte

func main() {
	field:=CreateField(7)

	p1:=CreatePoint(2,3,1)
	p2:=CreatePoint(5,3,1)
	p3:=CreatePoint(5,5,1)

	field.SetThreePoints(p1,p2,p3)
	field.PrintField()

	newP:=Create4Point(p1,p2,p3)
	fmt.Println(newP.x, newP.y)
	field.f[newP.x][newP.y]=1
	field.PrintField()




}
