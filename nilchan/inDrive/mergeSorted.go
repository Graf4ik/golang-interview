package inDrive

import (
	"fmt"
)

// Сложить два отсортированных массива

func main() {
	var a = []int{1, 5, 6, 18, 99}
	var b = []int{2, 4, 9, 11}

	fmt.Println(mergeSorted(a, b))
}

func mergeSorted(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	// добавляют остатки из a или b, если один из массивов закончился раньше другого.
	res = append(res, a[i:]...) // 🔹 Это значит: "вставь в res все элементы из first, начиная с индекса i"
	res = append(res, b[j:]...)

	return res
}

/*
Этот подход:

Использует знание, что a и b отсортированы.
Не требует дополнительной сортировки.
Работает за O(n + m) и без лишних копирований и нулей.
*/
