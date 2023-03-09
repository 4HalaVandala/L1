package main

import (
	"fmt"
	"strings"
)

/*
В данном примере мы создаем строку "snow dog sun — sun dog snow" и передаем ее в функцию reverseWords().
Функция разбивает строку на слова с помощью функции Split() и переворачивает каждое слово с помощью функции reverseString().
Затем она объединяет слова обратно в строку с помощью функции Join() и возвращает ее.
*/

func main() {
	str := "snow dog sun — sun dog snow"
	reversed := reverseWords(str)
	fmt.Println(reversed)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverseStrings(words[i])
	}
	return strings.Join(words, " ")
}

func reverseStrings(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
