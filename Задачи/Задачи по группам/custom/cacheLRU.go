package main

import "sync"

//===========================================================
//Задача 6
//1. Реализовать кеш. Для простоты считаем, что у нас бесконечная память и нам не нужно задумываться об удалении ключей из него.
//1. Почему использовал RWMutex, а не Mutex?
//2. Теперь представим что память не бесконечная. С какими проблемами столкнемся и как их решить?
//1. Какие есть алгоритмы выселения?
//3. Реализуй LRU
//===========================================================

// In-memory cache
// Нужно написать простую библиотеку in-memory cache.
// Реализация должна удовлетворять интерфейсу:

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Реализация кеша с бесконечной памятью (простой in-memory cache)

//type InMemoryCache struct {
//	data map[string]string
//	mu   sync.RWMutex
//}
//
//func (c *InMemoryCache) Set(k, v string) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	c.data[k] = v
//}
//
//func (c *InMemoryCache) Get(k string) (v string, ok bool) {
//	c.mu.RLock()
//	defer c.mu.Unlock()
//	v, ok = c.data[k]
//	return v, ok
//}

/*
❓ 2. Если память не бесконечная — с какими проблемами столкнёмся?
Память может переполниться, особенно при большом потоке Set.

Старые, давно неиспользуемые ключи могут навсегда оставаться в памяти, засоряя её.

✅ Решения:
Ввести ограничение на размер кеша.

Удалять старые элементы по определённому алгоритму выселения.
*/

/*
❓ 2.1. Какие есть алгоритмы выселения?
Наиболее популярные:

Алгоритм	Суть
LRU (Least Recently Used)	Удаляется наименее недавно использованный
LFU (Least Frequently Used)	Удаляется наименее часто используемый
FIFO (First In First Out)	Удаляется самый первый добавленный
Random Replacement	Удаляется случайный элемент
*/

type node struct {
	key, value string
	prev, next *node
}

type LRUCache struct {
	capacity   int
	mu         sync.Mutex
	items      map[string]*node
	head, tail *node // most recently used, least recently used
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*node, capacity),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, ok := c.items[key]; ok { // ok — true, если ключ существует, false — если нет.
		c.moveToFront(n)
		return n.value, true
	}

	/*
		Означает:
		"Если ключ key существует в кеше items, то переменной n присваивается
		 соответствующий узел (*node), и мы выполняем тело if".
		 Если ключа нет — ничего не делаем (ветка else или просто выходим).
	*/

	return "", false
}

func (c *LRUCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if n, ok := c.items[key]; ok {
		n.value = value
		c.moveToFront(n)
		return
	}

	if len(c.items) >= c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) moveToFront(n *node) {
	c.remove(n)
	c.addToFront(n)
}

func (c *LRUCache) addToFront(n *node) {
	n.prev = nil
	n.next = c.head
	if c.head != nil {
		c.head.prev = n
	}
	c.head = n
	if c.tail == nil {
		c.tail = n
	}
}

func (c *LRUCache) remove(n *node) {
	if n.prev == nil {
		n.prev.next = n.next
	} else {
		c.head.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		c.tail = n.prev
	}
}

func (c *LRUCache) removeOldest() {
	if c.tail != nil {
		delete(c.items, c.tail.key)
		c.remove(c.tail)
	}
}

/*
📌 Заключение
RWMutex нужен, когда много чтений и мало записей.
Если память ограничена — нужен алгоритм вытеснения.
Для наилучшего компромисса между скоростью и полезностью — LRU.
*/
