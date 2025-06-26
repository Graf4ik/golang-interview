package concurency

import (
	"fmt"
	"sync"
)

//===========================================================
//Задача 6
//1. Что выведет код и как исправить?
//===========================================================

/*
🧨 Проблемы
❗ Гонка по переменной a
Переменная a изменяется в трёх горутинах без синхронизации.
Это создаёт гонку данных, результат непредсказуем.
Значение a в финале (fmt.Printf("%d", a)) будет любое из трёх (10, 11 или 12), но может быть и повреждённое значение на уровне байт.
❗ Гонка по globalMap["test"] и другим слайсам
append на слайс не является потокобезопасной операцией.
Одновременное изменение []int может привести к ошибке выполнения или повреждению данных.
*/

var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
var a = 0
var mu sync.Mutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		mu.Lock()
		defer wg.Done()
		a = 10
		globalMap["test"] = append(globalMap["test"], a)
		mu.Unlock()

	}()
	go func() {
		mu.Lock()
		defer wg.Done()
		a = 11
		globalMap["test2"] = append(globalMap["test2"], a)
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		defer wg.Done()
		a = 12
		globalMap["test3"] = append(globalMap["test3"], a)
		mu.Unlock()
	}()
	wg.Wait()
	fmt.Printf("%v", globalMap)
	fmt.Printf("%d", a)
}
