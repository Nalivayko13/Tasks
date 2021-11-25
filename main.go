package main

import (
	"PointersTask2/geom"
	"fmt"
)


func main() {
	field:= geom.CreateField(8)

	p1:= geom.CreatePoint(3,3)
	p2:= geom.CreatePoint(4,3)
	p3:= geom.CreatePoint(5,5)
	p5:= geom.CreatePoint(4,2)

	pam:= geom.CreateParallelogram(p1,p2,p5,&field)
	field.PrintField()
	fmt.Println("P-am center is", pam.Center)
	fmt.Println("P-am points: ", pam.P1,pam.P2,pam.P3,pam.P4)
	fmt.Println("======================================")
	geom.ChangePoint(p1,p2,p3,&field)
	field.PrintField()




}
