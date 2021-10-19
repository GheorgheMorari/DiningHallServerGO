package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var tableStatus = [...]string{"Waiting for customers.", "Waiting for waiter.", "Waiting for delivery.", "Eating."}

type Table struct {
	id     int
	status int //Look at tableStatus
	order  *Order
}

func NewTable(id int, status int, order *Order) *Table {
	ret := new(Table)
	ret.id = id
	ret.status = status
	ret.order = order
	return ret
}

func (t *Table) deliver(delivery *Delivery, now int64) {
	if t.status != 2 {
		fmt.Printf("WRONG STATUS")
	}

	t.status = 3 //Set status to "eating"
	t.order = nil

	rating := 0
	maxWaitF := float64(delivery.MaxWait)
	timeWaitedF := float64(now - delivery.PickUpTime)
	if maxWaitF > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.1 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.2 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.3 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.4 > timeWaitedF {
		rating += 1
	}
	diningHall.ratings.addValue(rating)
	go func() {
		time.Sleep(timeUnit * (time.Duration(len(delivery.Items) + 1)))
		t.waitCustomers()
	}()
}

func (t *Table) waitCustomers() {
	t.status = 0
	t.order = generateOrder(t)

	time.Sleep(timeUnit * time.Duration(rand.Intn(10)))

	t.status = 1
}

func (t *Table) serve(waiter *Waiter) *Order {
	t.status = 2
	t.order.WaiterId = waiter.id
	t.order.PickUpTime = getUnixTimeUnits()
	return t.order
}

func (t *Table) getStatus() string {
	waitStatus := ""
	if t.order != nil && t.status == 2 {
		waitStatus = " Waiting for:" + strconv.Itoa(int(getUnixTimeUnits()-t.order.PickUpTime)) + "sec" + " Max wait:" + strconv.Itoa(t.order.MaxWait)
	}
	return "Table id:" + strconv.Itoa(t.id) + " Status:" + tableStatus[t.status] + waitStatus
}
