package tutu

import "fmt"

// отсортировать (руками)
func main() {
	slice := []string{"a", "b", "d", "c", "f", "e"}

	usort(slice)

	fmt.Printf("%v\n", slice)
}

// Сортировка пузырьком
func usort(slice []string) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice)-i-1; j++ {
			if slice[j] < slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
}

// С > → a b c d e f (по возрастанию — по алфавиту)
// С < → f e d c b a (по убыванию — в обратном порядке)
