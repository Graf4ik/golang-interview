package Магнит

import (
	"fmt"
	"sync"
	"time"
)

// Необходимо написать in-memory кэш

// 1. У кэша должен быть TTL
// должна быть возможность задать кастомный TTL для каждого эл-та
// 2. Написать функции  работы с кэшом:
// - Получение пользователя по его ID (при чтении не обновляем TTL)
// - Добавление в кеш
// - Удаление пользователя из кеша по его ID

type Cache struct {
	storage map[string]User
	times   map[string]time.Time // TTL на каждый элемент
	//TTL     time.Duration
	m *sync.RWMutex
}

type User struct {
	ID string
}

func (c *Cache) Add(user User, ttl time.Duration) error {
	c.m.Lock()
	c.storage[user.ID] = user
	c.times[user.ID] = time.Now().Add(ttl)
	c.m.Unlock()
	return nil
}

func (c *Cache) GetUserById(userId string) (*User, error) {
	c.m.RLock()
	t, ok := c.times[userId]
	if !ok || time.Now().After(t) {
		if time.Since(t) > c.TTL {
			c.m.Lock()
			delete(c.storage, userId)
			delete(c.times, userId)
			c.m.Unlock()
			return nil, fmt.Errorf("user not found")
		}
	}

	user, ok := c.storage[userId]
	c.m.RUnlock()
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (c *Cache) DeleteUser(userId string) (*User, error) {
	c.m.Lock()
	defer c.m.Unlock()
	user, ok := c.storage[userId]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	delete(c.storage, userId)
	delete(c.times, userId)
	return &user, nil
}

func main() {

}
