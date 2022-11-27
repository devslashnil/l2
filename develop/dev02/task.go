package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
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

var inputErr = errors.New("некорректный ввод")

func Unpack(str string) (string, error) {
	var b strings.Builder
	s := []rune(str)
	for i := 0; i < len(s); {
		// обработка невозможных событий как не обработанное число, слэш в конце
		if s[i] == '\\' && i+1 == len(s) || unicode.IsDigit(s[i]) {
			return "", inputErr
		} else if s[i] == '\\' {
			// при встрече слэша переходим к следующему символу
			i++
		}

		// если не дошли до конца просто записываем не численные константы в билдер
		// также если через символ число переключаемся на обработку такого кейса
		if i+1 == len(s) || !unicode.IsDigit(s[i+1]) {
			b.WriteRune(s[i])
			i++
		} else {
			// если через символ встретили число считываем его до конца и повторяем символ это число раз
			var digits []rune
			j := i + 1
			for j < len(s) && unicode.IsDigit(s[j]) {
				digits = append(digits, s[j])
				j++
			}
			count, _ := strconv.Atoi(string(digits))
			b.WriteString(strings.Repeat(string(s[i]), count))
			i = j
		}

	}

	return b.String(), nil

}

func main() {

}
