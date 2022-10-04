package cache

import (
	"L0/models"
	"L0/repository"
)

var cache map[string]models.Order

func Init() {
	cache = make(map[string]models.Order, 0)
}

func Set(order models.Order) {
	cache[order.OrderUID] = order
}

func GetByID(id string) (models.Order, bool) {
	order, ok := cache[id]
	return order, ok
}

func Recover() {
	for _, value := range repository.SelectOrders() {
		Set(value)
	}
}
