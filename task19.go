package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	var start int
	end := len(runes) - 1

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	return string(runes)
}

func main() {
	qwe := "stroka qwe"
	fmt.Println(reverseString(qwe))
}
