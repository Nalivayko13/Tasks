package main

type Field struct {
	n int64
}

func (f Field) CreateField(n int){
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
}

func (f *Field) SetThreePoints(p1, p2, p3 Point) {

}
