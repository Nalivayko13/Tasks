package main

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
	field:=CreateField(5)

	p1:=CreatePoint(4,2,1)
	p2:=CreatePoint(2,0,1)
	p3:=CreatePoint(1,3,1)

	field.SetThreePoints(p1,p2,p3)
	field.PrintField()


}
