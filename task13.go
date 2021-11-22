package main

import (
	"strings"
)

func ReverselineParentheses(s string) string{
	var Reverse = func(str string) string {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	var (
		start1, stop int
		out string
	)

	for i:=0;i<len(s);i++{
		start1=strings.Index(s,"(")
		stop=strings.Index(s,")")
		s=s[:start1]+Reverse(s[start1+1:stop])+s[stop+1:]
		out=s
		if !strings.Contains(s,"("){break}

	}

	return out
}


