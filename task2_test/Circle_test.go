package task2_test

import (
	"task2/geom"
	"testing"
)

func TestCreateCircle(t *testing.T){
	var pam geom.Parallelogram
	pam.P1.X=3
	pam.P1.Y=3
	pam.P2.X=4
	pam.P2.Y=3
	pam.P3.X=5
	pam.P3.Y=5
	pam.P4.X=4
	pam.P4.Y=2
	pam.Center.X=3
	pam.Center.Y=3
	pam.Square=2
	var expected geom.Circle
	expected.Center.Y=3
	expected.Center.X=3
	expected.Color=3
	expected.Square =2
	expected.Radius =0.7978845608028654
	expected.Length=5.013256549262001
	result:=geom.CreateCircle(pam)
	if result!=expected{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}

}

func TestCalcRadius(t *testing.T){
	var O geom.Circle
	O.Square=2.0
	expected:=0.7978845608028654
	O.CalcRadius(O.Square)
	result:=O.Radius
	if expected!=result{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}
}



