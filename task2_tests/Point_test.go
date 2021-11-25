package task2_tests

import (
	"PointersTask2/geom"
	"testing"
)

func TestCreatePoint(t *testing.T) {
	var expected geom.Point
	expected.X =2
	expected.Y =4
	expected.Color=3

	result:= geom.CreatePoint(2,4)

	if expected!=result{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}

}


