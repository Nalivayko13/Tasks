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

type Color int

func main() {
	field:=CreateField(7)

	p1:=CreatePoint(2,3,1)
	p2:=CreatePoint(5,3,1)
	p3:=CreatePoint(3,5,1)
	p5:=CreatePoint(4,2,2)

	pam:=CreateParallelogram(p1,p2,p5,field)
	field.PrintField()
	fmt.Println("P-am center is", pam.center)
	fmt.Println("P-am points: ", pam.p1,pam.p2,pam.p3,pam.p4)
	fmt.Println("======================================")
	ChangePoint(p1,p2,p3,field)
	field.PrintField()




}
