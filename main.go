package main

import (
	"fmt"
	"task2/geom"
)

func main() {
	field:= geom.CreateField(8)

	p1:= geom.CreatePoint(3,3)
	p2:= geom.CreatePoint(4,3)
	//p3:= geom.CreatePoint(5,5)
	p5:= geom.CreatePoint(4,2)

	pam:= geom.CreateParallelogram(p1,p2,p5,&field)
	field.PrintField()
	fmt.Println("P-am center is", pam.Center)
	fmt.Println("P-am points: ", pam.P1,pam.P2,pam.P3,pam.P4)
	fmt.Println("======================================")

	p6:=geom.CreatePoint(4,0)
	p7:=geom.CreatePoint(3,3)
	p8:=geom.CreatePoint(7,1)
	geom.ChangePoint(p6,p7,p8,&field)
	field.PrintField()




}

