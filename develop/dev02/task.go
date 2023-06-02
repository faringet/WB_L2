package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackS(s string) string {
	var scan scanner.Scanner
	var res string
	var prev string

	// "расчехляем" сканер
	scan.Init(strings.NewReader(s))

	// проставляем режим (на символы и инты)
	scan.Mode = scanner.ScanChars | scanner.ScanInts

	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		switch token {
		case scanner.Int:
			num, err := strconv.Atoi(scan.TokenText())
			if err != nil {
				fmt.Printf("Ошибка %s", err)
			}
			res += strings.Repeat(prev, num-1)
		default:
			prev = scan.TokenText()
			res += scan.TokenText()
		}
	}
	return res
}

func main() {
	fmt.Println(UnpackS("a4bc2d5e"))
	fmt.Println(UnpackS("abcd"))
	fmt.Println(UnpackS("45"))
	fmt.Println(UnpackS(""))

}
