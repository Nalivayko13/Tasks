package main

import "testing"

func TestCheckPalindrom(t *testing.T) {
	s:="aabaa"
	expected:=true

	result:=CheckPalindrom(s)

	if expected!=result {
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}
}
