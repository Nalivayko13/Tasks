package task2_test

import (
	"reflect"
	"task2/geom"
	"testing"
)

func TestCreteField(t *testing.T) {
	var expected geom.Field
	expected.F= [][]int{{0,0,0},{0,0,0},{0,0,0}}
	expected.N=3
	result:=geom.CreateField(3)
	flag:=reflect.DeepEqual(expected,result)
	if !flag {
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}

}

func TestCleanField(t *testing.T){
	var expected geom.Field
	expected.F= [][]int{{0,1,3},{0,6,12},{0,5,0}}
	expected.N=3
	expected.CleanField()
	flag:=true
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if expected.F[i][j]!=0{
				flag=false
			}
		}
	}
	if !flag{
		t.Errorf("incorrect")

	}

}
