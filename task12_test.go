package main

import "testing"

func TestMatrixElementsSum(t *testing.T) {
	n:=3
	expected:=6

	result:=MatrixElementsSum(n)

	if expected!=result {
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}
}
