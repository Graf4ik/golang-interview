package lofty

import (
	"fmt"
	"sync"
)

// Написать конкурентную реализацию. Дальше как можно масштабировать (партиции и тд)

//type MailLog interface {
//	Get(login string) (string, error)
//	Set(login string, mail string) error
//}

type MailLog interface {
	Get(login string) (string, error)
	Set(login string, mail string) error
}

type Mail struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMail() *Mail {
	return &Mail{
		data: make(map[string]string),
	}
}

func (m *Mail) Get(login string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	mail, ok := m.data[login]
	if !ok {
		return "", fmt.Errorf("login %s not exist", login)
	}
	return mail, nil
}

func (m *Mail) Set(login string, mail string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[login] = mail
	return nil
}

// Пример использования:
func main() {
	m := NewMail()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			login := fmt.Sprintf("user%d", i)
			m.Set(login, fmt.Sprintf("user%d@example.com", i))
		}(i)
	}

	wg.Wait()

	email, _ := m.Get("user3")
	fmt.Println("Email for user3:", email)
}

// 🚀 Масштабирование: шардирование (партиционирование)
//Когда записей становится много, и одна глобальная мьютекс-блокировка становится узким местом, можно распараллелить хранилище по N "шардам".
//
//🔸 Идея: использовать []*mailLog, выбирать shard по hash(login) % shardCount.
//go
//Копировать
//Редактировать
//type shardedMailLog struct {
//	shards []MailLog
//	count  int
//}
//
//func NewShardedMailLog(count int) MailLog {
//	shards := make([]MailLog, count)
//	for i := 0; i < count; i++ {
//		shards[i] = NewMailLog()
//	}
//	return &shardedMailLog{
//		shards: shards,
//		count:  count,
//	}
//}
//
//func (s *shardedMailLog) getShard(login string) MailLog {
//	h := fnvHash(login) % uint32(s.count)
//	return s.shards[h]
//}
//
//func (s *shardedMailLog) Get(login string) (string, error) {
//	return s.getShard(login).Get(login)
//}
//
//func (s *shardedMailLog) Set(login string, mail string) error {
//	return s.getShard(login).Set(login, mail)
//}
//
//func fnvHash(s string) uint32 {
//	var h uint32 = 2166136261
//	const prime32 = 16777619
//	for i := 0; i < len(s); i++ {
//		h *= prime32
//		h ^= uint32(s[i])
//	}
//	return h
//}
