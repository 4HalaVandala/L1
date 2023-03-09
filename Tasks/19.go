package main

import (
	"fmt"
	"strings"
)

/*
Для переворачивания строки в Go можно использовать пакет "strings". В нем есть функция Reverse, которая переворачивает переданную ей строку.
*/
func main() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Println(reversed)
}

func reverseString(s string) string {
	return strings.Join(reverseSlice(strings.Split(s, "")), "")
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
