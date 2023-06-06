package main

import (
	"fmt"
	"sort"
	"strings"
)

// SortWord сортируем символы в слове и приводим к нижнему регистру
func SortWord(s string) string {
	s = strings.ToLower(s)
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

// MakeUniqSlice получаем слайс только с уникальными значениями
func MakeUniqSlice(bucket []string) []string {
	// используем для удаления повторяющихся элементов
	set := make(map[string]bool)
	// используем для хранения уникальных элементов
	var setToSlice []string
	for _, v := range bucket {
		v = strings.ToLower(v)
		set[v] = true // присваиваем всегда true. Так как ключи должны быть уникальными, то повторяющиеся элементы будут удалены
	}
	for i := range set {
		// собираем уникальные элементы
		setToSlice = append(setToSlice, i)
	}
	sort.Strings(setToSlice)
	return setToSlice
}

// printAnagrams печатает в консоль
func printAnagrams(bucket map[string][]string) {
	for k, v := range bucket {
		fmt.Printf("%v : %v\n", k, v)
	}
}

// FinedAnagram является главной функцией, которая ищет во входном массиве анаграммы
func FinedAnagram(arr *[]string) map[string][]string {
	anag := make(map[string][]string, 0)
	anagV := make(map[string][]int, 0)
	res := make(map[string][]string, 0)
	anagKeys := make([]string, 0)

	for i, v := range *arr {
		sorted := SortWord(v)
		anagV[sorted] = append(anagV[sorted], i)
		anag[sorted] = append(anag[sorted], v)
	}

	for k, _ := range anag {
		anagKeys = append(anagKeys, k)
	}

	for _, v := range anagKeys {
		min := anagV[v][0]
		res[(*arr)[min]] = MakeUniqSlice(anag[v])
	}

	return res
}

func main() {
	words := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "ПяТаК", "ЛиСтОк", "тест"}
	anagrams := FinedAnagram(&words)
	printAnagrams(anagrams)
}
