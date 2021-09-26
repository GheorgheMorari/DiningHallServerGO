package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	Id      int `json:"id"`
	TableId int `json:"table_id"`
	WaiterId int   `json:"waiter_id"`
	Items    []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max_wait"`
	PickUpTime int64 `json:"pick_up_time"`
}
func (o *Order) getPayload()[]byte{
	result , err := json.Marshal(*o)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	return result
}
var orderIdCounter = 1

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
		Id: getOrderId(),
		//TODO configure table ids
		TableId: rand.Intn(10),
		//TODO configure waiter ids
		WaiterId:   rand.Intn(10),
		Items:      getRandomItems(),
		Priority:   rand.Intn(10),
		MaxWait:    rand.Intn(30)+20,
		PickUpTime: time.Now().Unix(),
	}
}
