package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var inputErr = errors.New("файл не поддерживает работу с данным флагом")

type Args struct {
	column            int  // k
	numeric           bool // n
	reverse           bool // r
	unique            bool // u
	month             bool // m
	trimEnd           bool // b
	check             bool // c
	numericWithSuffix bool // h
}

func ParseArgs() {

}

func PerformSort(args Args) {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось открыть файл: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var s []string
	for sc.Scan() {
		s = append(s, sc.Text())
	}

}

func Sort() {

}

func ColumnSort() {

}

func NumSort() {
}

func ReverseSort() {

}

func RemoveDublicates() {

}

func MonthSort() {

}

func TrimEndSort() {

}

func NumWithSuffixSort() {

}

func main() {

}
