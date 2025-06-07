package main

import "sync"

// ===========================================================
// Задача 7
// ===========================================================
// Для структуры

type SafeMap struct {
	data map[string]string
	mu   sync.RWMutex
}

func (m *SafeMap) GetOrCreate(key, value string) string {
	// Быстрая попытка чтения с RLock
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[key]

	if ok {
		return val
	}

	// Ключа нет — нужно заблокировать на запись
	m.mu.Lock()
	defer m.mu.Unlock()

	// Повторная проверка (double-check) — кто-то мог вставить значение пока мы ждали Lock
	if val, ok = m.data[key]; ok {
		return val
	}

	m.data[key] = value
	return value
}

// напишите потокобезопасный метод GetOrCreate(key, value string) string.
// Количество чтений на несколько порядков превышает количество вставок.
// Метод возвращает значение, которое уже есть в data или создает новое, если ключ не обнаружен
