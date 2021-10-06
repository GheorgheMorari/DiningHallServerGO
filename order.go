package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	Id         int   `json:"id"`
	TableId    int   `json:"table_id"`
	WaiterId   int   `json:"waiter_id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max_wait"`
	PickUpTime int64 `json:"pick_up_time"`
}

func (o *Order) getPayload() []byte {
	result, err := json.Marshal(*o)
	if err != nil {
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

//TODO add table and waiter
func generateOrder(table *Table) *Order {

	itemNum := rand.Intn(5) + 1
	var items []int
	maxWait := -1
	for i := 0; i < itemNum; i++ {
		item := rand.Intn(len(menu))
		items = append(items, item)
		itemWait := menu[item].preparationTime * 3
		if itemWait > maxWait {
			maxWait = itemWait
		}
	}
	ret := new(Order)

	ret.Id = getOrderId()
	ret.TableId = table.id
	ret.WaiterId = -1
	ret.Items = items
	ret.Priority = rand.Intn(3)
	ret.MaxWait = maxWait
	ret.PickUpTime = time.Now().Unix()

	return ret
}
