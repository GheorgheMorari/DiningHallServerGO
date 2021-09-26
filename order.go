package main

import (
	"math/rand"
	"time"
)

type Order struct {
	id         int
	tableId    int
	waiterId   int
	items      []int
	priority   int
	maxWait    int
	pickUpTime int64
}

var orderIdCounter = 0

func getOrderId() int {
	orderIdCounter++
	return orderIdCounter - 1
}
func getRandomItems() []int {
	var ret []int
	for i := 0; i < rand.Intn(10)+1; i++ {
		ret = append(ret, rand.Intn(10)+1)
	}
	return ret
}
func getRandomOrder() Order {
	return Order{
		id: orderIdCounter,
		//TODO configure table ids
		tableId: rand.Intn(10),
		//TODO configure waiter ids
		waiterId:   rand.Intn(10),
		items:      getRandomItems(),
		priority:   rand.Intn(10),
		maxWait:    rand.Intn(30)+20,
		pickUpTime: time.Now().Unix(),
	}
}
