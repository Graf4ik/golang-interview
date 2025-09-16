package МТС

func a(s []int) {
	s = append(s, 37) // 0 0 0 37
}
func b(m map[int]int) {
	m[3] = 33
}

// 1. Что будет содержать s после инициализации?
// 2. Что произойдет в println для слайса и для мапы?
func main() {
	s := make([]int, 3, 8) // 0 0 0
	m := make(map[int]int, 8)
	// add to slice
	a(s)
	println(s[3]) // ❗️ panic: index out of range
	// add to map
	b(m)
	println(m[3]) // 33
}

// 📌 3. println(s[3])
//❌ Паника: index out of range [3] with length 3
// ты не увеличивал len(s), ты просто делал append внутри a, не сохранив результат. s[3] не существует.
// s = make([]int, 3, 8) — длина 3, значит индексы 0,1,2. Доступ к s[3] вызывает панику.

// ✔ Как исправить:
// Если хочешь изменить исходный слайс — передавай указатель или возвращай значение:
// func a(s *[]int) {
//	*s = append(*s, 37)
// }
//...
// a(&s)
// fmt.Println(s[3]) // теперь доступен
