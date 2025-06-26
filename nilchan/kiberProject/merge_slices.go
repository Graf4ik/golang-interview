package kiberProject

import "fmt"

// смержить два слайса без дублей

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{2, 3, 4, 5, 6}

	mergeSlices(s1, s2)
	fmt.Println(mergeSlices(s1, s2))
}

// 1. Использование map для удаления дубликатов
func mergeSlices(s1, s2 []int) []int {
	unique := make(map[int]bool)
	res := []int{}

	// Сначала добавляем элементы из первого слайса
	for _, v := range s1 {
		if !unique[v] {
			unique[v] = true
			res = append(res, v)
		}
	}
	// Затем добавляем элементы из второго слайса, которых еще нет
	for _, v := range s2 {
		if !unique[v] {
			unique[v] = true
			res = append(res, v)
		}
	}

	return res
}
