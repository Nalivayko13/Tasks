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
	field:=CreateField(5)
	fmt.Println("hello")
	field.PrintField()

}
