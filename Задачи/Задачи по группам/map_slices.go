package main

// =========================================
// Задача 1
// Что выведет код?
// =========================================

//func main() {
//	v := []int{3, 4, 1, 2, 5}
//	ap(v)
//	sr(v)
//	fmt.Println(v) // 1, 2, 3, 4, 5 +
//}
//
//func ap(arr []int) {
//	arr = append(arr, 10)
//}
//
//func sr(arr []int) {
//	sort.Ints(arr)
//}

//==========================================
//Задача 2
//1. Что выведет код?
//==========================================

//func main() {
//	var foo []int
//	var bar []int
//
//	foo = append(foo, 1) // foo = [1] cap=1
//	foo = append(foo, 2) // foo = [1 2] cap=2
//	foo = append(foo, 3) // foo [1 2 3], cap=4 (возможно, Go увеличил запас)
//	bar = append(foo, 4) // bar = [1 2 3 4] (на том же backing array, если cap(foo) ≥ 4)
//	foo = append(foo, 5) // foo: [1 2 3 5], но и bar может измениться, если они все еще делят массив
//
//	fmt.Println(foo, bar) // [1 2 3 5] [1 2 3 5]
//}

// ===========================================================
// Задача 3
// 1. Что выведется?
// ===========================================================
//func main() {
//	c := []string{"A", "B", "D", "E"}
//	b := c[1:2]
//	//	Memory (shared backing array):
//	//	Index:   0    1    2    3
//	//	c:     [A]  [B]  [D]  [E]
//	//				 ↑
//	//			b starts here
//	//
//	//	b := c[1:2] → b = ["B"], len=1, cap=3
//	b = append(b, "TT")
//	//	Memory (same backing array gets modified):
//	//	Index:   0    1     2     3
//	//	c:     [A]  [B]  [TT]  [E]
//	//				 ↑
//	//		b now = ["B", "TT"]
//	fmt.Println(c) // ["A", "B", "TT", "E"] потому что c[2] теперь "TT"
//	fmt.Println(b) // ["B", "TT"]
//}

// ===========================================================
// Задача 4
// 1. Что выведет код?
// ===========================================================

//func main() {
//	var m map[string]int
//	// 	m := make(map[string]int) Чтобы всё работало нужно инициализировать мапу
//	for _, word := range []string{"hello", "world", "from", "the",
//		"best", "language", "in", "the", "world"} {
//		m[word]++
//	}
//	for k, v := range m {
//		fmt.Println(k, v) // panic: assignment to entry in nil map
//	}
//}

//===========================================================
//Задача 5
//1. Что будет в результате выполнения?
//===========================================================

//func main() {
//	// внутри mutate вы работаете с копией структуры слайса, но она указывает на тот же массив.
//	mutate := func(a []int) {
//		a[0] = 0
//		/*  0, 2, 3, 4 Вы меняете первый элемент в массиве — это видно и снаружи,
//		потому что a и аргумент в mutate указывают на один и тот же массив. */
//		a = append(a, 1)
//		fmt.Println(a) // 0 2 3 4 1
//	}
//	a := []int{1, 2, 3, 4}
//	mutate(a)
//	fmt.Println(a) // 0, 2, 3, 4
//	/*
//		Внутри функции: [0 2 3 4 1]
//		Снаружи: [0 2 3 4] ← изменён первый элемент, но append не повлиял
//	*/
//}

//==========================================================
//Задача 6
//1. Что выведется?
//2. Зная обо всех таких нюансах, которые могут возникнуть, какие есть рекомендации?
//===========================================================

// # Вариант 1
// -----------
//func mod(a []int) {
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // [5 5 5 5] +
//}
//
//func main() {
//	sl := []int{1, 2, 3, 5}
//	mod(sl)
//	fmt.Println(sl) // [5 5 5 5] +
//}

// # Вариант 2
// -----------
//func mod(a []int) {
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // 5 5 5 5 +
//}
//func main() {
//	sl := make([]int, 4, 8)
//	sl[0] = 1
//	sl[1] = 2
//	sl[2] = 3
//	sl[3] = 5
//	mod(sl)
//	fmt.Println(sl) // 5 5 5 5 +
//}

// # Вариант 3
// -----------
//func mod(a []int) {
//	a = append(a, 125)
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // 5 5 5 5 5 +
//}
//func main() {
//	sl := make([]int, 4, 8)
//	sl[0] = 1
//	sl[1] = 2
//	sl[2] = 3
//	sl[3] = 5
//	mod(sl)
//	fmt.Println(sl) // 5 5 5 5 +
//}

// # Вариант 4
// -----------
//func mod(a []int) {
//	a = append(a, 125)
//	for i := range a {
//		a[i] = 5
//	}
//	fmt.Println(a) // 5 5 5 5 5 5 +
//}
//func main() {
//	sl := []int{1, 2, 3, 4, 5}
//	mod(sl)
//	fmt.Println(sl) // 1 2 3 4 5 +
//}

// ===========================================================
// Задача 7
// 1. Что будет содержать s после инициализации?
// 2. Что произойдет в println для слайса и для мапы?
// ===========================================================
//func a(s []int) {
//	s = append(s, 37) // добавляет 37, но в КОПИЮ слайса!
//}
//
//func b(m map[int]int) {
//	m[3] = 33
//}
//
//func main() {
//	s := make([]int, 3, 8)
//	m := make(map[int]int, 8)
//
//	// add to slice
//	a(s)
//	// println(s[3]) // ✅ компилируется, т.к. cap(s) = 8
//	// Но len(s) всё ещё 3, и доступ к s[3] выходит за границы длины
//	// panic: runtime error: index out of range [3] with length 3.
//
//	// add to map
//	b(m)
//	println(m[3]) // 33
//}

//===========================================================
//Задача 8
//1. Расскажи подробно что происходит
//===========================================================
//# Вариант 1
//-----------

//func main() {
//	a := []int{1, 2}
//	a = append(a, 3)  // 1 2 3 (cap 4)
//	b := append(a, 4) // a сейчас: [1, 2, 3] (len=3, cap=4), b указывает на этот новый слайс
//	c := append(a, 5) // a всё тот же: [1, 2, 3] Добавляем 5 → результат: [1, 2, 3, 5]
//
//	fmt.Println(b) // 1 2 3 5
//	fmt.Println(c) // 1 2 3 5
//}

//# Вариант 2
//-----------

//func main() {
//	a := []int{1, 2}
//	a = append(a, 3)  // 1 2 3 cap = 4
//	a = append(a, 7)  // 1 2 3 7 cap = 4
//	b := append(a, 4) // 1 2 3 7 4 (новый массив)
//	c := append(a, 5) // 1 2 3 7 5 (новый массив)
//
//	fmt.Println(b) // 1 2 3 7 4
//	fmt.Println(c) // 1 2 3 7 5
//}

// ===========================================================
// Задача 9
// Что выведет код и почему?
// ===========================================================
//func main() {
//	foo := make([]int, 0, 4)
//	foo = append(foo, 1, 2, 3) // 1 2 3 (cap=4)
//	bar := append(foo, 4)      // 1 2 3 5 (тот же массив)
//	baz := append(foo, 5)      // 1 2 3 5 (тот же массив)
//
//	fmt.Println(bar) // 1 2 3 5
//	fmt.Println(baz) // 1 2 3 5
//}
