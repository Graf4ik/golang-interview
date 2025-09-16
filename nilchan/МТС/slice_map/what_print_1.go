package МТС

import (
	"fmt"
	"sort"
)

// Что выведет код?
func main() {
	v := []int{3, 4, 1, 2, 5} // 5 5
	ap(v)
	sr(v)
	fmt.Println(v) // [1 2 3 4 5]
}
func ap(arr []int) {
	arr = append(arr, 10) // 5 5 10
}
func sr(arr []int) {
	sort.Ints(arr) //  сортирует срез по ссылке
}

// 🧠 Почему не [3 4 1 2 5]?
// Потому что sort.Ints(v) изменяет содержимое среза,
// так как срез передаётся по ссылке (но append — нет, если результат не возвращается).
