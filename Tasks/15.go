package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}

Приведет к утечке памяти т.к. "justString" будет хранить ссылку на большу строку до зовершения программы
Для решения этой проблемы можно скопировать результаты с помощью функции copy()
Память будет очищена послев выполнения someFunc()
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	copyString := make([]byte, 100)
	copy(copyString, v)
	justString = string(copyString)
}

func main() {
	someFunc()
}
