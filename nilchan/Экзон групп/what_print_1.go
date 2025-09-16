package main

import "fmt"

// какие тут ошибки, что выведет?
// После замены Sprintf на fmt.Println(elm), как работает Sprintf и попросил указать fmt.Printf("%s;", elm)
// Будет ли %s работать с экземпляром структур?
// Как в Го реализовать универсальное приведение структуры к строке?
type MyContainer struct {
	int A
}

func main() {
	arr := []MyContainer{
		{A: 1}, {A: 2}, {A: 3},
	}
	modify(arr)
	for _, elm := range arr {
		fmt.Sprintf("%s;", elm)
		// ❌ Ошибка компиляции:
		// fmt.Sprintf("%s;", elm)
		// %s ожидает строку, а elm — это структура MyContainer, а не string.
		// Go не знает, как отобразить структуру как строку с помощью %s.
		// 👉 Это вызовет panic: reflect: call of reflect.Value.String on struct Value, если принудительно исполнить.
	}
}

func modify(arr []MyContainer) {
	arr[0].A = 10
	arr = append(arr, MyContainer{A: 4}, MyContainer{A: 5})
}

// ✅ 2. Что выведет этот код как есть?
// Ничего.
// fmt.Sprintf("%s;", elm) создаёт строку, но она не печатается — Sprintf возвращает строку, но её никто не использует (результат игнорируется).

// ✅ 3. Что произойдёт, если заменить Sprintf на fmt.Println(elm)?
// for _, elm := range arr {
//	fmt.Println(elm)
// }
// Вывод будет: {10} {2} {3}
// Почему:
// modify(arr):
// arr[0].A = 10 → изменит первый элемент.
// append(...) не изменит оригинальный слайс (принимается по значению).
// В main() остался только [ {10}, {2}, {3} ].

// ✅ 4. Как работает Sprintf?
// s := fmt.Sprintf("%s", val)
// Sprintf — возвращает строку (в отличие от Printf, который печатает её).
// %s — для строк.
// %v — для универсального форматирования (включая структуры).
// %+v — покажет имена полей.
// %#v — покажет Go-литерал (удобно для отладки).
// ✅ Пример:
// elm := MyContainer{A: 10}
// fmt.Printf("%v\n", elm)   // {10}
// fmt.Printf("%+v\n", elm)  // {A:10}
// fmt.Printf("%#v\n", elm)  // main.MyContainer{A:10}

// ✅ 5. Можно ли использовать %s со структурой?
// Нет, если только структура не реализует fmt.Stringer интерфейс, т.е. метод:
// func (m MyContainer) String() string {
//	return fmt.Sprintf("MyContainer: %d", m.A)
// }
// Тогда %s будет работать, пример:
// type MyContainer struct {
//	A int
// }
// func (m MyContainer) String() string {
//	return fmt.Sprintf("A=%d", m.A)
// }
// func main() {
//	m := MyContainer{A: 42}
//	fmt.Printf("%s\n", m) // A=42
// }

// ✅ 6. Как реализовать универсальное приведение структуры к строке в Go?
// 👉 Реализовать метод String() для структуры:
//
// type Stringer interface {
//	String() string
// }
// Пример:
// type User struct {
//	Name string
//	Age  int
// }
// func (u User) String() string {
//	return fmt.Sprintf("User{Name: %s, Age: %d}", u.Name, u.Age)
// }

// ✅ Вывод по коду:
// Ошибки:
// %s нельзя использовать с MyContainer, пока не реализован String().
// Sprintf(...) ничего не делает — результат не используется.
// Рекомендации:
// Используй %v для структур, если не нужен свой String().
// Для красивого вывода — реализуй String().
// append() в modify() не изменит исходный слайс — если нужно изменить, возвращай результат или передавай *[]MyContainer.
