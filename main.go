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
	p3:=CreatePoint(3,5,1)
	p5:=CreatePoint(4,2,2)

	field.SetThreePoints(p1,p2,p5)
	field.PrintField()

	//p4:=Create4Point(p1,p2,p3, field)
	//field.PrintField()
	//var circle Circle
	//circle.CalcRadius(CalcSquare(p2,p3,p4))
	fmt.Println("======================")
	pam:=CreateParallelogram(p1,p2,p3,field)
	fmt.Println(pam.center)
	fmt.Println("**********************************")
	ChangePoint(p1,p2,p5,field)
	field.PrintField()




}
