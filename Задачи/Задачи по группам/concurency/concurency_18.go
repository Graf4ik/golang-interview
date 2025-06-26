package concurency

import (
	"sync"
)

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
func balance() (int, error) { // 25 sec
	x := make(map[int]int, 1)
	var m sync.Mutex

	wg := sync.WaitGroup{}
	errCh := make(chan error, 5)

	// call bank
	for i := 0; i < 5; i++ {
		//i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			m.Lock()
			b, err := bank_network_call(num)
			if err != nil {
				errCh <- err
				return
			}

			x[num] = b
			m.Unlock()
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		return 0, err // Возвращаем первую ошибку
	}

	sum := 0
	for _, v := range x {
		sum += v
	}
	// Как-то считается сумма значений в мапе и возвращается
	return sum, nil
}

// Вариант 2 (c errgroup)

func balance2() (int, error) {
	x := make(map[int]int, 5)
	var m sync.Mutex
	g := new(errgroup.Group)

	// Запускаем 5 горутин через errgroup
	for i := 0; i < 5; i++ {
		i := i // Фиксируем значение i для горутины
		g.Go(func() error {
			b, err := bank_network_call(i)
			if err != nil {
				return err // Группа автоматически прервется при первой ошибке
			}

			m.Lock()
			x[i] = b
			m.Unlock()
			return nil
		})
	}

	// Ждем завершения всех горутин. Если была ошибка — возвращаем её.
	if err := g.Wait(); err != nil {
		return 0, err
	}

	// Считаем сумму
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, nil
}
