package main

import "testing"

func TestReverselineParentheses(t *testing.T) {
	s:="foo(bar)baz"
	expected:="foorabbaz"

	result:=ReverselineParentheses(s)

	if expected!=result {
		t.Errorf("incorrect. Expected %v, got %v",expected,result)
	}
}
