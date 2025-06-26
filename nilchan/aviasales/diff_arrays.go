package aviasales

// написать функцию, которая вернет разновидность двух массивов
// (массив, в котором собраны все элементы первого, но нет пар во втором массиве)

// а если есть повторы, еще нужно сохранить порядок

func main() {
	diff([]int{1, 2, 3}, []int{1, 2})
	diff([]int{1, 2, 1, 3}, []int{3})
}
func diff(left, right []int) []int {
	// Создаём мапу для подсчёта количества каждого элемента во втором массиве
	rightCount := make(map[int]int)

	for _, v := range right {
		rightCount[v]++
	}

	result := make([]int, 0, len(left))
	for _, v := range left {
		if rightCount[v] > 0 {
			rightCount[v]-- // "используем" один экземпляр
			continue
		}
		result = append(result, v)
	}

	return result
}
