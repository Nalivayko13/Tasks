package main

func CheckPalindrom(s string) bool{
	for i:=0;i<len(s)/2;i++{
		if s[i]!=s[len(s)-i-1]{
			return false
		}
	}
	return true
}



func main() {
	//fmt.Println(CheckPalindrom("aabaa"))
	//fmt.Println(CheckPalindrom("aabba"))
	//fmt.Println(MatrixElementsSum(3))
	//fmt.Println(ReverselineParentheses("foo(bar)baz"))
}
