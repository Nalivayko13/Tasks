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
	X, Y int
	Color
}

func CreatePoint(x,y int) Point {
	var p Point
	p.X =x
	p.Y =y
	p.Color= Color((p.X + p.Y) / 2)
	return p
}


