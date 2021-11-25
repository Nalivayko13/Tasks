package task2_tests

import (
	"PointersTask2/geom"
	"testing"
)

func TestCreteField(t *testing.T) {
	var expected geom.Field
	expected.F= [][]int{{0,0,0},{0,0,0},{0,0,0}}
	expected.N=3
	result:=geom.CreateField(3)
	flag:=true
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if expected.F[i][j]!=result.F[i][j]{
				flag=false
			}

		}
	}

	if !flag {
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}

}

func TestPrintField(t *testing.T){
	var expected geom.Field
	expected.F= [][]int{{0,0,0},{0,0,0},{0,0,0}}
	expected.N=3
}