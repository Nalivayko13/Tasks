package geom


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

type Point struct {
	x,y int
	Color
}

func CreatePoint(x,y int) Point {
	var p Point
	p.x=x
	p.y=y
	p.Color= Color((p.x + p.y) / 2)
	return p
}

