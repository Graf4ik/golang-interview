package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// нужно оставить только уникальные
type Order struct {
	ID     string `json:"id"`     // ID заказа (поле в JSON)
	Amount int    `json:"amount"` // Сумма заказа (поле в JSON)
}

func main() {
	orders, err := loadOrders("orders.json")
	if err != nil {
		fmt.Println("Error loading orders:", err)
		return
	}
	unique := make(map[string]bool)
	uniqueOrders := make([]Order, 0, len(orders))

	for _, order := range orders {
		if !unique[order.ID] {
			unique[order.ID] = true
			uniqueOrders = append(uniqueOrders, order)
		}
	}

	// Обрабатываем заказы (выводит их в консоль)
	processOrders(uniqueOrders)
}

func loadOrders(filename string) ([]Order, error) {
	// Читаем содержимое файла
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var orders []Order
	// Десериализуем JSON  в массив структур Order
	err = json.Unmarshal(data, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func processOrders(orders []Order) {
	for _, order := range orders {
		fmt.Println("Processing loading order:", order.ID)
	}
}
