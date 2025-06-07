package leetcode

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Сортируем интервалы по началу
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	// Добавляем первый интервал в результат
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1] // последнй интервал в res
		current := intervals[i]

		if current[0] <= last[1] { // сравнивается первый и последний элемент предыдущего интервала
			// Объединяем с предыдущим
			res[len(res)-1][1] = max(last[1], current[1])
		} else {
			// Добавляем новый интервал
			res = append(res, current)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	fmt.Println(merged) // [[1 6] [8 10] [15 18]]
}

/*
Пример: [[1,3],[2,6],[8,10],[15,18]]
Сортируем: [[1,3],[2,6],[8,10],[15,18]] (уже отсортированы)

Стартуем с [1,3]

[2,6] пересекается (2 ≤ 3) → объединяем → [1,6]

[8,10] не пересекается → добавляем

[15,18] не пересекается → добавляем

Результат: [[1,6],[8,10],[15,18]]
*/
