package Ситидрайв

// 1. что выведет

func a(s []int) {
	s = append(s, 37)
}

func b(m map[int]int) {
	m[3] = 33
}

func main() {
	s := make([]int, 3, 8) // 0 0 0
	m := make(map[int]int, 8)

	a(s)
	println(s[3]) // panic: runtime error: index out of range [3] with length 3

	b(m)
	println(m[3]) // 33
}
