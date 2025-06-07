package livecoding_checklist

import (
	"fmt"
	"sort"
)

// 19. Сортировка слайса встроенной функцией.
// В Go для сортировки слайсов используется пакет sort,
// который предоставляет несколько удобных функций для разных типов данных.

func main() {
	// 1. Сортировка слайса чисел (int, float64)
	// Сортировка целых чисел
	nums := []int{4, 2, 7, 1, 5}
	sort.Ints(nums)
	fmt.Println(nums) // [1 2 4 5 7]

	// Сортировка чисел с плавающей точкой
	floats := []float64{3.2, 1.5, 4.7, 2.1}
	sort.Float64s(floats)
	fmt.Println(floats) // [1.5 2.1 3.2 4.7]

	// 2. Сортировка строк
	names := []string{"Иван", "Анна", "Петр", "Мария"}
	sort.Strings(names)
	fmt.Println(names) // [Анна Иван Мария Петр]

	// Сортировка в обратном порядке
	nums2 := []int{4, 2, 7, 1, 5}

	// Сортировка в обратном порядке
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	fmt.Println(nums2) // [7 5 4 2 1]

	// Сортировка пользовательских структур
	// 1. Реализация интерфейса sort.Interface
	people := []Person3{
		{"Иван", 30},
		{"Анна", 25},
		{"Петр", 40},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people) // [{Анна 25} {Иван 30} {Петр 40}]

	// 2. Использование sort.Slice (более простой способ)
	people2 := []Person3{
		{"Иван", 30},
		{"Анна", 25},
		{"Петр", 40},
	}

	// Сортировка по возрасту (возрастание)
	sort.Slice(people2, func(i, j int) bool {
		return people2[i].Age < people2[j].Age
	})
	fmt.Println(people2) // [{Анна 25} {Иван 30} {Петр 40}]

	// Сортировка по имени (убывание)
	sort.Slice(people2, func(i, j int) bool {
		return people2[i].Name > people2[j].Name
	})
	fmt.Println(people2) // [{Петр 40} {Иван 30} {Анна 25}]

	// Стабильная сортировка (сохранение порядка равных элементов)
	people3 := []Person3{
		{"Иван", 30},
		{"Анна", 25},
		{"Мария", 30},
		{"Петр", 40},
	}

	// Стабильная сортировка по возрасту
	sort.SliceStable(people3, func(i, j int) bool {
		return people3[i].Age < people3[j].Age
	})
	fmt.Println(people3)
	// Порядок "Иван" и "Мария" сохранится относительно исходного
	// [{Анна 25} {Иван 30} {Мария 30} {Петр 40}]

	// Проверка отсортированности
	nums3 := []int{1, 2, 3, 4, 5}
	fmt.Println(sort.IntsAreSorted(nums3)) // true

	names2 := []string{"Борис", "Анна", "Виктор"}
	fmt.Println(sort.StringsAreSorted(names2)) // false
}

type Person3 struct {
	Name string
	Age  int
}

// ByAge реализует sort.Interface для сортировки []Person по возрасту
type ByAge []Person3

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

/*
Полезные советы

Для простых типов используйте sort.Ints, sort.Float64s, sort.Strings
Для пользовательских структур предпочтительнее sort.Slice (короче и проще)
Если важен порядок равных элементов - используйте sort.SliceStable
Для сортировки в обратном порядке применяйте sort.Reverse
Для очень больших слайсов рассмотрите возможность использования более специализированных алгоритмов сортировки
*/
