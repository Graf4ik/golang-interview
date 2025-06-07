package livecoding_checklist

import "fmt"

// 4. Уметь инстанциировать массив, слайс и мапу, из литералов и через make.
//То же самое для разных типов данных в них, в том числе структур, в том числе пустых структур.

func main() {
	// Пустой массив int
	arrInt := [3]int{} // [0 0 0]

	// Массив int с значениями
	arrIntInit := [3]int{1, 2, 3} // [1 2 3]

	// Массив структур
	type Point struct{ X, Y int }
	arrStruct := [2]Point{
		{1, 2},
		{3, 4},
	}

	// Массив пустых структур
	arrEmptyStruct := [2]struct{}{{}, {}}

	// Массивы нельзя создать через make(), так как их размер фиксирован
	// Используйте литералы или new()
	arr := new([3]int) // &[0 0 0]

	// Пустой слайс int
	sliceInt := []int{} // []

	// Слайс int с значениями
	sliceIntInit := []int{1, 2, 3} // [1 2 3]

	// Слайс структур
	sliceStruct := []Point{
		{1, 2},
		{3, 4},
	}

	// Слайс пустых структур
	sliceEmptyStruct := []struct{}{{}, {}}

	// Слайс int с длиной 3
	sliceInt := make([]int, 3) // [0 0 0]

	// Слайс int с длиной 3 и вместимостью 5
	sliceIntCap := make([]int, 3, 5) // [0 0 0]

	// Слайс структур
	sliceStruct := make([]Point, 2) // [{0 0} {0 0}]

	// Слайс пустых структур
	sliceEmptyStruct := make([]struct{}, 2) // [{} {}]

	// Пустая мапа string -> int
	mapStrInt := map[string]int{} // map[]

	// Мапа с значениями
	mapStrIntInit := map[string]int{
		"one": 1,
		"two": 2,
	}

	// Мапа структур
	mapStrStruct := map[string]Point{
		"first":  {1, 2},
		"second": {3, 4},
	}

	// Мапа пустых структур
	mapStrEmptyStruct := map[string]struct{}{
		"key1": {},
		"key2": {},
	}

	// Мапа string -> int
	mapStrInt := make(map[string]int) // map[]

	// Мапа string -> int с начальной вместимостью
	mapStrIntCap := make(map[string]int, 10) // map[]

	// Мапа структур
	mapStrStruct := make(map[string]Point) // map[]

	// Мапа пустых структур
	mapStrEmptyStruct := make(map[string]struct{}) // map[]

	/*
		Особенности:
		Массивы имеют фиксированный размер, указанный при объявлении
		Слайсы динамические, могут расти с помощью append
		Мапы динамические, автоматически растут при добавлении элементов
		Пустые структуры struct{} не занимают память (размер 0)
		При использовании make() для слайсов можно указать длину и вместимость
		Для мап make() позволяет указать начальную вместимость (но не размер)
	*/
	// Использование как множества (set)
	set := make(map[string]struct{})
	set["item1"] = struct{}{}
	set["item2"] = struct{}{}

	// Проверка наличия
	if _, exists := set["item1"]; exists {
		fmt.Println("item1 exists")
	}

	// Канал пустых структур для сигналов
	signalChan := make(chan struct{})
}
