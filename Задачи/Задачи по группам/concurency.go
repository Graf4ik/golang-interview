package main

import (
	"context"
)

//==========================================
//Задача 1
//Что выведет код? Исправить все проблемы
//==========================================

/*
❌ Проблемы:
Канал ch не закрывается — range ch ожидает закрытия канала, но main никогда его не закрывает → бесконечное ожидание.
Горутины могут блокироваться при ch <- v*v, потому что канал небуферизованный, а main не читает из него, пока все wg.Done() не сработают. Это создаёт потенциальный deadlock.
Канал читается после wg.Wait(), но горутины могут быть уже заблокированы при попытке записи в ch.
*/

//func main() {
//	ch := make(chan int, 3) // Буферизированный канал
//	wg := &sync.WaitGroup{}
//	wg.Add(3)
//	for i := 0; i < 3; i++ {
//		go func(v int) {
//			defer wg.Done()
//			ch <- v * v
//		}(i)
//	}
//	go func() {
//		wg.Wait()
//		defer close(ch)  // Закрываем канал после завершения всех горутин
//	}()
//
//	var sum int
//	for v := range ch {
//		sum += v
//	}
//	fmt.Printf("result: %d\n", sum)
//}

//==========================================
//Задача 2
//Что выведет код? Должны выводится все значения
//==========================================

/*
В main() нет ожидания завершения горутин, поэтому программа завершится сразу после запуска всех go fmt.Println(i),
не дождавшись их выполнения. В результате:
✅ Иногда вы увидите часть значений от 0 до 4999.
❌ Но чаще — очень мало значений или ничего вообще, потому что main() завершает выполнение раньше, чем горутины успеют отработать.
*/

//func main() {
//	a := 5000
//
//	wg := sync.WaitGroup{}
//	wg.Add(a)
//
//	for i := 0; i < a; i++ {
//		go func(i int) {
//			defer wg.Done()
//			fmt.Println(i)
//		}(i)
//	}
//	wg.Wait()
//}

// ===========================================================
// Задача 3
// Будет ошибка что все горутины заблокированы. Какие горутины будут заблокированы? И почему?
// ===========================================================

/*
 Важное замечание:
ch <- 1 — блокирует main, потому что канал небуферизированный и никто ещё не слушает его.
А go func(...) — даже не начнёт выполняться, потому что main уже заблокирован и не дойдёт до создания горутины.
В итоге единственная активная горутина (main) уже стоит на блокирующей операции, а других ещё не существует.
*/
//func main() {
//	ch := make(chan int)
//	// ch <- 1 // ➋ main блокируется здесь, ожидая, что кто-то примет значение из канала
//	go func() { // ➌ эта горутина НИКОГДА НЕ ЗАПУСТИТСЯ
//		fmt.Println(<-ch)
//	}()
//
//	ch <- 1
//}

//===========================================================
//Задача 4
//1. Как это работает, что не так, что поправить?
//===========================================================

/*
❗ Почему возникает deadlock?
Небуферизированный канал блокирует send (ch <- true), пока нет активного получателя.
А вы пишете в канал до запуска горутины, которая могла бы читать.
В результате: main-горутина зависает, а горутина даже не запускается.
Go фиксирует это и вызывает панику
*/
//func main() {
//	ch := make(chan bool)
//
//	ch <- true // (2) ПЫТАЕМСЯ отправить true — но НИКТО не читает!
//	// 💥 DEADLOCK на этой строке: main заблокировался, а горутина ещё не создана!
//
//	go func() {
//		<-ch // (3) Горутина будет читать — но до неё дело не дошло
//	}()
//	ch <- true

// Вариант 1
// ch := make(chan bool)

//go func() {
//	<-ch
//	<-ch
//}()
//
//ch <- true
//ch <- true

// Вариант 2
//ch := make(chan bool, 2) // буфер на 2 значения
//
//ch <- true  // записываем — НЕ блокируется
//ch <- true  // тоже записываем
//
//go func() {
//	<-ch
//	<-ch
//}()
//}

//===========================================================
//Задача 5
//1. Как будет работать код?
//2. Как сделать так, чтобы выводился только первый ch?
//===========================================================

//func main() {
//	ch := make(chan bool)
//	ch2 := make(chan bool)
//	ch3 := make(chan bool)
//
//	// Вариант 1: запускать только первую горутину (остальные не запускать)
//	// Вариант 2: задержать остальные горутины с помощью time.Sleep
//	go func() {
//		time.Sleep(100 * time.Millisecond)
//		ch <- true
//	}()
//	go func() {
//		time.Sleep(100 * time.Millisecond)
//		ch2 <- true
//	}()
//	go func() {
//		ch3 <- true
//	}()
//
//	/*
//	select выбирает случайный case, если несколько каналов готовы одновременно.
//	Поэтому результат будет недетерминированным: на каждом запуске может вывестись:
//	*/
//	select {
//	case <-ch:
//		fmt.Printf("val from ch")
//	case <-ch2:
//		fmt.Printf("val from ch2")
//	case <-ch3:
//		fmt.Printf("val from ch3")
//	}
//}

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
//var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
//var a = 0
//var mu sync.Mutex
//
//func main() {
//	wg := sync.WaitGroup{}
//	wg.Add(3)
//	go func() {
//		mu.Lock()
//		defer wg.Done()
//		a = 10
//		globalMap["test"] = append(globalMap["test"], a)
//		mu.Unlock()
//
//	}()
//	go func() {
//		mu.Lock()
//		defer wg.Done()
//		a = 11
//		globalMap["test2"] = append(globalMap["test2"], a)
//		mu.Unlock()
//	}()
//	go func() {
//		mu.Lock()
//		defer wg.Done()
//		a = 12
//		globalMap["test3"] = append(globalMap["test3"], a)
//		mu.Unlock()
//	}()
//	wg.Wait()
//	fmt.Printf("%v", globalMap)
//	fmt.Printf("%d", a)
//}

//===========================================================
//Задача 7
//===========================================================

type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

/*
Пояснение:
Все функции запускаются параллельно.
Как только одна из них успешно возвращает результат — он отдается вызывающему коду.
Ошибки накапливаются, последняя возвращается, если все SearchFunc завершились с ошибками.
ctx можно использовать для отмены всех запросов (если передан извне, например с таймаутом).
Дополнительно можно убрать time.After(...), если не нужно ограничивать по времени помимо ctx.
*/
//func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
//
//	// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
//	// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
//	// с ошибкой - отдаем последнюю полученную ошибку
//
//	errCh := make(chan error, len(sfs))
//	results := make(chan Result, len(sfs))
//
//	wg := sync.WaitGroup{}
//	wg.Add(len(sfs))
//
//	for _, sf := range sfs {
//		go func(sf SearchFunc) {
//			defer wg.Done()
//
//			res, err := sf(ctx, query)
//			if err != nil {
//				errCh <- err
//				return
//			}
//
//			// отправляем результат, но только если канал открыт
//			select {
//			case results <- res:
//			default:
//			}
//		}(sf)
//	}
//
//	// Закрываем каналы, когда все горутины завершатся
//	go func() {
//		wg.Wait()
//		defer close(errCh)
//		defer close(results)
//	}()
//
//	// ждем первый результат или все ошибки
//	select {
//	case res := <-results:
//		return res, nil
//	case <-ctx.Done():
//		return Result{}, ctx.Err()
//	case <-time.After(100 * time.Millisecond): // safety net timeout, опционально
//	}
//
//	// Все вернули ошибку — возвращаем последнюю
//	var lastErr error
//	for err := range errCh {
//		lastErr = err
//	}
//
//	return Result{}, lastErr
//}

//===========================================================
//Задача 8
//1. Что выведется и как исправить?
//===========================================================

/*
Текущий код содержит две основные ошибки:
Нет ожидания завершения горутин.
Нет синхронизации при доступе к counter.
Либо через mutex либо через atomic
*/

//func main() {
//	// var counter int
//	var counter int32
//	wg := sync.WaitGroup{}
//	// mu := sync.Mutex{}
//
//	//for i := 0; i < 1000; i++ {
//	//	wg.Add(1)
//	//	go func() {
//	//		defer wg.Done()
//	//		mu.Lock()
//	//		counter++
//	//		mu.Unlock()
//	//	}()
//	//}
//
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			atomic.AddInt32(&counter, 1)
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println(counter)
//}

//===========================================================
//Задача 9
//1. Что выведется и как исправить?
//2. Что поправить, чтобы сохранить порядок?
//===========================================================

/*
Проблема с переменной i: Передаем i как параметр в анонимную функцию
Это решает проблему захвата переменной в замыкании
Синхронизация: Добавлен sync.WaitGroup для ожидания завершения всех горутин
Закрываем канал после завершения всех отправок
Сохранение порядка: Используем буферизированный канал как очередь
Выводим сообщения в том же потоке (main), сохраняя порядок
Другие улучшения: Добавлен defer для гарантированного вызова wg.Done()
Используем range для чтения из канала
*/
// Вариант 1
//func main() {
//	m := make(chan string, 5) // Буфер на все сообщения
//	cnt := 5
//	for i := 0; i < cnt; i++ {
//		go func() {
//			m <- fmt.Sprintf("Goroutine %d", i)
//		}()
//	}
//	for i := 0; i < cnt; i++ {
//		fmt.Println(<-m)
//	}
//}

// Вариант 2
//	func main() {
//		m := make(chan string, 3)
//		cnt := 5
//
//		var wg sync.WaitGroup
//
//		for i := 0; i < cnt; i++ {
//			wg.Add(1)
//			go func(id int) {
//				defer wg.Done()
//				m <- fmt.Sprintf("Goroutine %d", id)
//			}(i) // Передаем i как параметр
//		}
//  	// Получение (в отдельной горутине, чтобы не блокировать main)
//		go func() {
//			wg.Wait()
//			close(m)
//		}()
//
//		for msg := range m {
//			fmt.Println(msg)
//		}

//    	// Вывод в порядке отправки (используем буферизированный канал)
//		//for i := 0; i < cnt; i++ {
//		//	go ReceiveFromCh(m)
//		//}
//	}
//func ReceiveFromCh(ch chan string) {
//	fmt.Println(<-ch)
//}

//===========================================================
//Задача 10
//1. Merge n channels
//2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы
//===========================================================

//func case3(channels ...chan int) chan int {
//	out := make(chan int)
//
//	var wg sync.WaitGroup
//	wg.Add(len(channels))
//
//	for _, ch := range channels {
//		go func() {
//			defer wg.Done()
//			for v := range ch {
//				out <- v
//			}
//		}()
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//
//		for _, c := range channels {
//			select {
//			case _, ok := <-c:
//				if ok {
//					close(c)
//				}
//			default:
//				close(c)
//			}
//		}
//	}()
//
//	return out
//}

//func mergeChannelsWithCancel(channels ...chan int) chan int {
//	out := make(chan int)
//	ctx, cancel := context.WithCancel(context.Background())
//	var wg sync.WaitGroup
//
//	forward := func(c chan int) {
//		defer wg.Done()
//		for {
//			select {
//			case num, ok := <-c:
//				if !ok {
//					cancel() // Инициируем отмену при закрытии любого канала
//					return
//				}
//				out <- num
//			case <-ctx.Done():
//				return
//			}
//		}
//	}
//
//	wg.Add(len(channels))
//	for _, c := range channels {
//		go forward(c)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//		cancel()
//		for _, c := range channels {
//			close(c)
//		}
//	}()
//
//	return out
//}

//===========================================================
//Задача 11
//	1. Описать словами. Предположим есть метод REST API. В нем мы хотим сделать 10 запросов к другим API.
//	Нужно считать данные  и отправить пользователю. Как это сделать? Как добавить таймаут?
//	Стоит ли использовать каналы или можно WaitGroup?
//===========================================================

//===========================================================
//Задача 12
//1. Конурентно по батчам запросить данные и записать в файл. Нужна общая конструкция, функции которые делают запрос к сайту и выгрузку в файл можно не реализовывать.
//2. Сделать так, чтобы одновременно выполнялось не более chunkSize запросов.
//===========================================================
/*
 Пояснение:
sem := make(chan struct{}, chunkSize) — ограничивает количество одновременных горутин (и, соответственно, запросов).
worker — запускается для каждого id, делает запрос и сохраняет результат.
wg — синхронизация завершения всех воркеров.
fetchData и saveData — заглушки, которые ты можешь заменить на реальную логику.

*/
// Константы
//const (
//	urlTemplate = "http://jsonplaceholder.typicode.com/tools/%d"
//	chunkSize   = 100
//	dataCount   = 2 << 10 // 2048
//)
//
//// Заглушка для запроса (эмуляция)
//func fetchData(id int) (string, error) {
//	// Здесь можно сделать реальный HTTP-запрос, например:
//	// resp, err := http.Get(fmt.Sprintf(urlTemplate, id))
//	// defer resp.Body.Close()
//	// Но пока просто возвращаем заглушку:
//	return fmt.Sprintf("Data for ID %d", id), nil
//}
//
//// Заглушка для записи в файл (можно сделать буфер и потом слить в файл)
//func saveData(id int, data string) error {
//	// Здесь можно реализовать запись в файл или канал
//	fmt.Printf("Saving ID %d: %s\n", id, data)
//	return nil
//}
//
//func worker(id int, wg *sync.WaitGroup, sem chan struct{}) {
//	defer wg.Done()
//
//	// Ограничение количества одновременно работающих горутин
//	sem <- struct{}{}        // блокируем
//	defer func() { <-sem }() // разблокируем
//
//	// Запрос данных
//	data, err := fetchData(id)
//	if err != nil {
//		fmt.Printf("Error fetching ID %d: %v\n", id, err)
//		return
//	}
//
//	// Сохранение
//	err = saveData(id, data)
//	if err != nil {
//		fmt.Printf("Error saving ID %d: %v\n", id, err)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	sem := make(chan struct{}, chunkSize) // семафор на chunkSize одновременных запросов
//
//	for i := 1; i <= dataCount; i++ {
//		wg.Add(1)
//		go worker(i, &wg, sem)
//	}
//
//	wg.Wait()
//	fmt.Println("All done.")
//}

//===========================================================
//Задача 13
//1. Запросить параллельно данные из источников. Если все где-то произошла ошибка, то вернуть ошибку, иначе вернуть nil.
//2. Представим, что теперь функция должна возвращать результат int. Есть функция resp.Size(), для каждого url
//надо проссумировать и вернуть, если ошибок не было. Просто описать подход к решению
//3. Что делать, если урлов у нас миллионы?
//
//===========================================================

//func main() {
//	_, err := download([]string{
//		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
//		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
//		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
//		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
//		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
//	})
//
//	if err != nil {
//		panic(err)
//	}
//}
//
//const maxWorkers = 100
/*
Параллельно запросить данные из источников. Если все запросы вернули ошибку — вернуть ошибку, иначе nil.
Подход:
Запустить параллельные запросы (через горутины).
Использовать sync.WaitGroup для ожидания завершения всех.
Считать количество ошибок, желательно через atomic.
Если все запросы завершились с ошибкой — вернуть ошибку.
*/

/*
Теперь download() должна возвращать int — сумму resp.Size(), если не было ошибок.
Подход:
Предполагаем, что resp.Size() — это размер тела ответа (например, len(body)).
Создаём totalSize int32/int64, безопасно увеличиваем его (atomic.AddInt64()).
Если какая-либо ошибка — возвращаем её.
Если все с ошибкой — вернуть ошибку, иначе вернуть сумму.
*/

//func download(urls []string) (int64, error) {
//	var wg sync.WaitGroup
//	var errorCount int32
//	var totalSize int64
//
//	for _, url := range urls {
//		wg.Add(1)
//		go func(url string) {
//			defer wg.Done()
//			resp, err := http.Get(url)
//			if err != nil || resp.StatusCode != http.StatusOK {
//				atomic.AddInt32(&errorCount, 1)
//				return
//			}
//			defer resp.Body.Close()
//
//			size := resp.ContentLength
//			if size > 0 {
//				atomic.AddInt64(&totalSize, size)
//			}
//		}(url)
//	}
//
//	wg.Wait()
//
//	if int(errorCount) == len(urls) {
//		return 0, errors.New("all requests failed")
//	}
//
//	return totalSize, nil
//}

/*
Что делать, если URL-ов миллионы?
Подход:
Ограничение параллелизма:
Создать семафор (например, workerPool := make(chan struct{}, maxWorkers)), чтобы не запускать миллионы горутин.
Обработка по чанкам или через очередь:
Использовать worker pool + канал задач, чтобы эффективно и стабильно обрабатывать огромный список URL-ов.
Возможен подход с batching + временными файлами, если нужно сохранять данные (например, response body) и не хватает оперативной памяти.
*/

//func download(urls []string) (int64, error) {
//	wg := sync.WaitGroup{}
//	wg.Add(len(urls))
//
//	var totalSize int64
//	var errorCount int32
//
//	jobs := make(chan string, len(urls))
//
//	// Воркеры
//	for i := 0; i < maxWorkers; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for url := range jobs {
//				resp, err := http.Get(url)
//				if err != nil || resp.StatusCode != http.StatusOK {
//					atomic.AddInt32(&errorCount, 1)
//					continue
//				}
//				size := resp.ContentLength
//				if size > 0 {
//					atomic.AddInt64(&totalSize, size)
//				}
//				resp.Body.Close()
//			}
//		}()
//	}
//
//	wg.Wait()
//
//	for _, url := range urls {
//		jobs <- url
//	}
//
//	close(jobs)
//
//	if int(errorCount) == len(urls) {
//		return 0, errors.New("all requests failed")
//	}
//
//	return totalSize, nil
//}

//===========================================================
//Задача 14
//1. Что выведет на экран и сколько времени будет работать?
//2. Нужно ускорить, чтобы работало быстрее. Сколько будет работать теперь?
//3. Если бы в networkRequest выполнялся реальный сетевой вызов, то какие с какими проблемами мы могли бы столкнуться в данном коде?
//4. Если url немного, а запросов к ним много, то как можно оптимизировать?
//===========================================================

/*
1. Каждый запрос занимает ~1 мс (time.Sleep(time.Millisecond))
Запросы выполняются параллельно (горутины)
Основное ограничение - количество CPU и contention на мьютексе
Ориентировочное время: 100-500 мс (зависит от CPU)
*/

/*
2. Как ускорить?
Проблемы текущей реализации:
Мьютекс создает contention при инкременте счетчика
Слишком много горутин (10k) создают overhead
Улучшения:
atomic.Int32 вместо мьютекса
Семафор для ограничения параллелизма (100 горутин)
Уменьшение накладных расходов на создание горутин
*/

/*
3. Проблемы с реальными сетевыми запросами
Основные проблемы:
Нет таймаутов:
Зависший запрос заблокирует горутину навсегда
Решение: context.WithTimeout
Нет обработки ошибок:
Ошибки сети просто игнорируются
Решение: возвращать ошибки из networkRequest
Нет retry логики:
Временные сбои приведут к потере запросов
Решение: добавить retry с экспоненциальным backoff
Нет ограничения скорости:
Можно перегрузить сервер
Решение: rate limiting (например, golang.org/x/time/rate)
Нет circuit breaker:
При проблемах с сервером продолжит посылать запросы
Решение: добавить паттерн Circuit Breaker
*/

/*
Если URL повторяются:
Кеширование ответов:
var cache = sync.Map{}
func networkRequest(url string) {
	if val, ok := cache.Load(url); ok {
		return val
	}
	// ... выполняем запрос
	cache.Store(url, result)
}
Пул соединений:
Использовать http.Client с настроенным Transport
Поддерживать keep-alive соединения
Пакетные запросы:
Объединять одинаковые запросы в batch
Пример: вместо 100 запросов к /user/1, сделать 1 запрос к /users?ids=1,2,3...
Шардирование запросов:
var urlShards = make([]chan string, 10)
// Каждая горутина обрабатывает свой shard
*/

//const numRequests = 10000
//
//var count atomic.Int32 // Атомарный счетчик
//// var count int
//var cache sync.Map
//var m sync.Mutex
//
//func networkRequest() {
//	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
//	m.Lock()
//	count++
//	m.Unlock()
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(numRequests)
//
//	sem := make(chan struct{}, 100) // Ограничиваем параллелизм
//	urls := generateUrls()          // []string с повторениями
//
//	for _, url := range urls {
//		sem <- struct{}{}
//		go func(u string) {
//			defer func() { <-sem }()
//
//			if _, ok := cache.Load(u); ok {
//				return
//			}
//
//			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//			defer cancel()
//
//			if err := realNetworkRequest(ctx, u); err == nil {
//				cache.Store(u, true)
//				count.Add(1)
//			}
//		}(url)
//	}

// Вариант 1
//for i := 0; i < numRequests; i++ {
//	sem <- struct{}{}
//	go func() {
//		defer func() {
//			wg.Done()
//			<-sem
//		}()
//		networkRequest()
//	}()
//}

//wg.Wait()
//fmt.Println(count)
//}

//===========================================================
//Задача 15
//
//===========================================================

// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).

// Нужно написать обертку predictableFunc,
// которая будет работать с заданным фиксированным таймаутом (например, 1 секунду).

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

// Есть функция, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
//func unpredictableFunc() int64 {
//	rnd := rand.Int63n(5000)
//	time.Sleep(time.Duration(rnd) * time.Millisecond)
//
//	return rnd
//}

// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.

// Вариант 1
//func predictableFunc() (int64, error) {
//	resultChan := make(chan int64, 1)
//	start := time.Now()
//	go func() {
//		resultChan <- unpredictableFunc()
//	}()
//
//	select {
//	case result := <-resultChan:
//		elapsed := time.Since(start)
//		fmt.Printf("Function executed in %v\n", elapsed)
//		return result, nil
//	case <-time.After(1 * time.Second):
//		return 0, fmt.Errorf("timeout after 1 second")
//	}
//}

// Вариант 2
//func predictableFunc() (int64, error) {
//	// Создаем контекст с таймаутом 1 секунда
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel() // Освобождаем ресурсы контекста
//
//	// Канал для результата
//	resultChan := make(chan int64, 1)
//
//	// Запускаем функцию в горутине, передавая контекст
//	start := time.Now()
//	go func() {
//		resultChan <- unpredictableFunc()
//	}()
//
//	// Ожидаем результат или отмену контекста
//	select {
//	case result := <-resultChan:
//		elapsed := time.Since(start)
//		fmt.Printf("Function executed in %v\n", elapsed)
//		return result, nil
//	case <-ctx.Done():
//		return 0, ctx.Err() // Возвращаем ошибку таймаута
//	}
//}

//func main() {
//	fmt.Println("started")
//
//	if result, err := predictableFunc(); err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Result:", result)
//	}
//}

//===========================================================
//Задача 16
//===========================================================

// Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.

//func merge(cs ...<-chan int) <-chan int {
//	out := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(len(cs))
//
//	for _, c := range cs {
//		go func() {
//			defer wg.Done()
//			for ch := range c {
//				out <- ch
//			}
//		}()
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}

//===========================================================
//Задача 17
//1. Что выведется? Исправь проблему
//===========================================================

// # Вариант1
// ----------
//func main() {
//	x := make(map[int]int, 1)
//	mu := sync.Mutex{}
//
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 7
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 10
//	}()
//	time.Sleep(100 * time.Millisecond)
//	fmt.Println("x[1] =", x[1])
//}

//===========================================================
//Задача 18
//1. Иногда приходят нули. В чем проблема? Исправь ее
//2. Если функция bank_network_call выполняется 5 секунд, то за сколько выполнится balance()? Как исправить проблему?
//3. Представим, что bank_network_call возвращает ошибку дополнительно. Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.
//===========================================================

/*
Пояснения:
1. Почему раньше приходили нули?
Горутины работали асинхронно, и balance() мог завершиться до того, как они запишут данные в x.
Исправлено: Добавлен sync.WaitGroup, который блокирует balance() до завершения всех горутин.
2. Сколько теперь выполняется balance(), если bank_network_call занимает 5 секунд?
Все 5 вызовов bank_network_call выполняются параллельно, поэтому balance() завершится за ~5 секунд (а не за 25, если бы запросы шли последовательно).
3. Как обрабатываются ошибки?
Если хотя бы один bank_network_call вернет ошибку, она попадет в errCh.

После wg.Wait() проверяем канал ошибок. Если есть хотя бы одна ошибка — возвращаем её.
Если ошибок нет — считаем сумму.
Дополнительные улучшения:
Буферизированный канал ошибок (errCh):
Размер буфера = 5, чтобы горутины не блокировались при отправке ошибок.
defer wg.Done():
Гарантирует, что WaitGroup уменьшит счетчик даже при панике.
close(errCh) после wg.Wait():
Позволяет безопасно использовать range для проверки ошибок.
*/

// Вариант 1
//func balance() (int, error) { // 25 sec
//	x := make(map[int]int, 1)
//	var m sync.Mutex
//
//	wg := sync.WaitGroup{}
//	errCh := make(chan error, 5)
//
//	// call bank
//	for i := 0; i < 5; i++ {
//		//i := i
//		wg.Add(1)
//		go func(num int) {
//			defer wg.Done()
//			m.Lock()
//			b, err := bank_network_call(num)
//			if err != nil {
//				errCh <- err
//				return
//			}
//
//			x[num] = b
//			m.Unlock()
//		}(i)
//	}
//
//	wg.Wait()
//	close(errCh)
//
//	for err := range errCh {
//		return 0, err // Возвращаем первую ошибку
//	}
//
//	sum := 0
//	for _, v := range x {
//		sum += v
//	}
//	// Как-то считается сумма значений в мапе и возвращается
//	return sum, nil
//}

// Вариант 2 (c errgroup)

//func balance() (int, error) {
//	x := make(map[int]int, 5)
//	var m sync.Mutex
//	g := new(errgroup.Group)
//
//	// Запускаем 5 горутин через errgroup
//	for i := 0; i < 5; i++ {
//		i := i // Фиксируем значение i для горутины
//		g.Go(func() error {
//			b, err := bank_network_call(i)
//			if err != nil {
//				return err // Группа автоматически прервется при первой ошибке
//			}
//
//			m.Lock()
//			x[i] = b
//			m.Unlock()
//			return nil
//		})
//	}
//
//	// Ждем завершения всех горутин. Если была ошибка — возвращаем её.
//	if err := g.Wait(); err != nil {
//		return 0, err
//	}
//
//	// Считаем сумму
//	sum := 0
//	for _, v := range x {
//		sum += v
//	}
//	return sum, nil
//}

//===========================================================
//Задача 19
//Что выведет код и почему?
//===========================================================

/*
Проблемы в коде:
Гонка данных (Data Race)
Переменная ch доступна из двух горутин (main и анонимной) без синхронизации.
Это неопределенное поведение: компилятор/процессор могут оптимизировать код так, что изменения ch в одной горутине не будут видны в другой.
runtime.GOMAXPROCS(1)
При использовании только одного ядра планировщик Go может вообще не переключаться между горутинами, пока main не заблокируется (например, на системном вызове).
В данном случае main занят активным циклом (for ch == 0), поэтому горутина ch = 1 может никогда не выполниться.
Активный цикл (busy loop)
Цикл for ch == 0 не освобождает процессор, поэтому планировщик Go не передает управление другой горутине.
Что выведет программа?

Вариант 1: Бесконечный цикл (гоутина ch = 1 не получает управление).
Вариант 2 (теоретически): Завершится с выводом "finish", если горутина ch = 1 все же выполнится (но это маловероятно при GOMAXPROCS(1)).
*/

//func main() {
//	runtime.GOMAXPROCS(1)
//	ch := 0
//	go func() {
//		ch = 1
//	}()
//	for ch == 0 {
//	}
//	fmt.Println("finish")
//}
