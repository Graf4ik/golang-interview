package Точка_банк

type Account struct {
	Id      int
	Name    string
	Balance float32
}

// реализация скрыта, но там что-то вроде sqlx
type AccountRepository interface {
	Find(id int) (*Account, error)
	Save(acc *Account) error
}

// в реализации отправка в какой-то брокер, но может и быть http post
type EventSender interface {
	Send(event interface{}) error
}

type TransferEvent struct {
	FromId int
	ToId   int
	Amount float32
}

type AccountService struct {
	repository  AccountRepository
	eventSender EventSender
}

// вызывается из rest-контроллера
func (s *AccountService) Transfer(fromId int, toId int, amount float32) error {
	fromAcc, err := s.repository.Find(fromId)
	if err != nil {
		return err
	}
	toAcc, err := s.repository.Find(toId)
	if err != nil {
		return err
	}

	toAcc.Balance += amount
	fromAcc.Balance -= amount

	err = s.repository.Save(fromAcc)
	if err != nil {
		return err
	}
	err = s.repository.Save(toAcc)
	if err != nil {
		return err
	}
	go func() {
		s.eventSender.Send(&TransferEvent{
			FromId: fromId,
			ToId:   toId,
			Amount: amount,
		})
	}()
	return nil
}

// 🔍 1. Нарушение атомарности и отсутствует транзакция
// Проблема:
// Вы сначала изменяете fromAcc и toAcc, затем вызываете Save дважды.
// Между двумя вызовами Save() может произойти ошибка, и данные будут несогласованны (например, деньги списались, но не зачислились).
// 📌 Решение:
// Использовать транзакцию в AccountRepository (если это SQL).
// Ввести TransferTx или WithTx(func(r AccountRepository) error).
//
// 🔍 2. Потенциальные гонки при параллельных переводах
// Проблема:
// Если два перевода происходят одновременно с одним и тем же аккаунтом, баланс может быть перезаписан неправильно (classic race condition при read-modify-write).
//
// 📌 Решение:
//
// На уровне базы: SELECT ... FOR UPDATE для блокировки строк.
// Или блокировка в Go: sync.Mutex на Account.Id (можно с sync.Map и map[int]*sync.Mutex).
//
// 🔍 3. Асинхронная отправка событий — без обработки ошибок
// Проблема:
// Вы вызываете go s.eventSender.Send(...), но не обрабатываете ошибку, не логируете, не перезапускаете, не подтверждаете.
// 📌 Решение:
// Использовать очередь/буфер событий с retry, логами и алертами.
// Или в случае не-критичных событий хотя бы логировать ошибку:
// go func() {
//    if err := s.eventSender.Send(...); err != nil {
//        log.Printf("failed to send event: %v", err)
//    }
// }()
// 🔍 4. Нейминг: Save — неясно, обновляет ли или вставляет
// 📌 Рекомендация:
// Сделать явные методы Update(*Account) и Insert(*Account), если возможно.

// 🔍 5. Float32 для денег — потенциальная ошибка округления
// Проблема:
// float32 (и float64) не подходят для представления денег из-за потерь точности при арифметике.
// 📌 Решение:
// Использовать int64 и хранить сумму в копейках (центах).
// Или тип decimal.Decimal (например, shopspring/decimal).
//
// 🔍 6. Плохая читаемость: длинный метод без разбивки
// 📌 Рекомендация:
// Разбить метод на подметоды:
// func (s *AccountService) Transfer(...) error {
//    fromAcc, toAcc, err := s.loadAccounts(fromId, toId)
//    ...
//    s.updateBalances(...)
//    ...
//    s.sendEventAsync(...)
// }
