package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// в структуре храним значения флагов
type inputFlags struct {
	column      int
	sortByNum   bool
	reverseSort bool
	unique      bool
	filename    []string
}

// парсим аргументы юзера (и заполняем inputFlags), которые он ввел при запуске
func parsArguments() inputFlags {
	columnNumber := flag.Int("k", 1, "указание колонки для сортировки")
	num := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	flags := inputFlags{
		column:      *columnNumber,
		sortByNum:   *num,
		reverseSort: *reverse,
		unique:      *unique,
		filename:    flag.Args(),
	}
	return flags
}

// ParseFile парсим файл и его имя (путь до которого указал юзер при запуске) и пишем в массив
func ParseFile(inputData string) ([]string, error) {
	var data []string
	file, err := os.Open(inputData)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return data, nil
}

// пишем в массив все данные из файла
func readFromFile(inputData []string) ([]string, error) {
	var data []string
	for _, v := range inputData {
		docData, err := ParseFile(v)
		if err != nil {
			return []string{}, err
		}
		data = append(data, docData...)
	}

	return data, nil
}

// Unique отбираем уникальные элементы и пишем их в новый массив
func Unique(data []string) []string {
	set := make(map[string]bool)
	var setToSlice []string
	for _, v := range data {
		set[v] = true
	}

	for i := range set {
		setToSlice = append(setToSlice, i)
	}

	sort.Strings(setToSlice)

	return setToSlice
}

// Sort выполняем сортировку в соответствии с флагами
func Sort(data []string, flags inputFlags) []string {
	if flags.unique {
		data = Unique(data)
	}

	compareAsNumbers := func(i, j string) bool {
		lnum, lerr := strconv.Atoi(i)
		rnum, rerr := strconv.Atoi(j)

		if lerr != nil && rerr != nil {
			return i < j
		}
		if lerr != nil || rerr != nil {
			return lerr == nil
		}
		return lnum < rnum
	}
	compareAsStrings := func(lhs, rhs string) bool {
		return lhs < rhs
	}

	var valueComparator func(string, string) bool
	if flags.sortByNum {
		valueComparator = compareAsNumbers
	} else {
		valueComparator = compareAsStrings
	}

	compareLogic := func(i, j int) bool {
		lhs := strings.Split(data[i], " ")
		rhs := strings.Split(data[j], " ")
		if len(lhs) == 0 {
			return true
		}
		if len(rhs) == 0 {
			return false
		}

		if len(lhs) < flags.column && len(rhs) >= flags.column {
			return true
		}
		if len(lhs) >= flags.column && len(rhs) < flags.column {
			return false
		}

		if len(lhs) < flags.column && len(rhs) < flags.column {
			return valueComparator(lhs[0], rhs[0])
		}
		if len(lhs) >= flags.column && len(rhs) >= flags.column {
			return valueComparator(lhs[flags.column-1], rhs[flags.column-1])
		}
		panic("заглушка")
	}

	if !flags.reverseSort {
		sort.Slice(data, compareLogic)
	} else {
		sort.Slice(data, func(i, j int) bool {
			return !compareLogic(i, j)
		})
	}

	return data
}

func main() {
	par := parsArguments()
	data, err := readFromFile(par.filename)
	if err != nil {
		log.Fatalln(err)
	}

	data = Sort(data, par)

	for _, line := range data {
		fmt.Println(line)
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, str := range data {
		_, err := f.WriteString(str + " ")
		if err != nil {
			log.Fatal(err)
		}
	}
}
