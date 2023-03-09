package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)
	charSet := make(map[rune]bool)
	for _, char := range str {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))                    // true
	fmt.Println(isUnique("abCdefAaf"))               // false
	fmt.Println(isUnique("aabcd"))                   // false
	fmt.Println(isUnique("go"))                      // true
	fmt.Println(isUnique("Go programming language")) // false
}
