package task2_test

import (
	"reflect"
	"task2/geom"
	"testing"
)
func TestCreateParallelogram(t *testing.T){
	var expected geom.Parallelogram
	expected.P1.X=4
	expected.P1.Y=0
	expected.P1.Color=2
	expected.P2.X=3
	expected.P2.Y=3
	expected.P2.Color=3
	expected.P3.X=7
	expected.P3.Y=1
	expected.P3.Color=4
	expected.P4.X=6
	expected.P4.Y=4
	expected.P4.Color=5
	expected.Center.X=3
	expected.Center.Y=1
	expected.Center.Color=2
	expected.Square=10

	var f geom.Field
	f.N=8
	f.F= make([][]int, 8)
	for i := range f.F {
		f.F[i] = make([]int, 8)
	}
	for i:=0;i<8;i++{
		for j:=0;j<8;j++{
			f.F[i][j]=0
		}
	}

	p1:=geom.Point{4,0, 2}
	p2:=geom.Point{3,3, 3}
	p3:=geom.Point{7,1, 4}
	result:=geom.CreateParallelogram(p1,p2,p3,&f)

	flag:=reflect.DeepEqual(result,expected)
	if !flag{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}

}

func TestCreate4Point(t *testing.T){
	expected:=geom.Point{6,4,5}

	var f geom.Field
	f.N=8
	f.F= make([][]int, 8)
	for i := range f.F {
		f.F[i] = make([]int, 8)
	}
	for i:=0;i<8;i++{
		for j:=0;j<8;j++{
			f.F[i][j]=0
		}
	}

	p1:=geom.Point{4,0, 2}
	p2:=geom.Point{3,3, 3}
	p3:=geom.Point{7,1, 4}

	result:=geom.Create4Point(p1,p2,p3,&f)
	if expected!=result{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}
}

func TestFindCenter(t *testing.T){
	expected:=geom.Point{3,1,2}
	p1:=geom.Point{4,0, 2}
	p2:=geom.Point{3,3, 3}
	flag1,flag2:=true, true
	result:=geom.FindCenter(p1,p2,&flag1,&flag2)

	if expected!=result{
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}



}