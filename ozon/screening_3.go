package main

import "fmt"

// дано два слайса, нужно объединить попарно в один
// Вариант 1
func main2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	res := make([][2]int, len(s1)) // Тип [][2]int означает: слайс, каждый элемент которого — массив из 2 чисел.

	for i := range s1 {
		res[i] = [2]int{s1[i], s2[i]}
		/*
		res[0] = [2]int{1, 4}
		res[1] = [2]int{2, 5}
		res[2] = [2]int{3, 6}
		*/
	}
	fmt.Println(res)
}

// Вариант 2
// произвольное количество слайсов одинаковой длины

func main() {
	// произвольное количество слайсов одинаковой длины
	//slices := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	slices := [][]int{}

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{7, 8, 9}

	slices = append(slices, s1)
	slices = append(slices, s2)
	slices = append(slices, s3)

	if len(slices) == 0 {
		return
	}

	length := len(slices[0])
	result := make([][]int, length)

	for i := 0; i < length; i++ {
		group := make([]int, len(slices)) // новый слайс group длиной len(slices)
		for j, s := range slices {
			group[j] = s[i]
			/*
				Берём i-й элемент из каждого слайса и помещаем его в j-ю позицию group.
				🔎 Пример:
				На первой итерации (i == 0):
				group[0] = slices[0][0] == 1
				group[1] = slices[1][0] == 4
				group[2] = slices[2][0] == 7
				=> group = []int{1, 4, 7}
			*/
		}
		result[i] = group // Сохраняем сформированную group на i-ю позицию результата.
	}

	fmt.Println(result) // [[1 4 7] [2 5 8] [3 6 9]]
}
