package main

import (
	"strconv"
	"time"
)

var waiterStatus = [...]string{"Waiting.", "Sending order id:", "Delivering delivery id:", "Waiting for orderList to clear."}

type Waiter struct {
	id         int
	atWork     int
	statusId   int
	modifierId int
}

func NewWaiter(id int, atWork int, statusId int, modifierId int) *Waiter {
	ret := new(Waiter)
	ret.id = id
	ret.atWork = atWork
	ret.statusId = statusId
	ret.modifierId = modifierId
	return ret
}

func (w *Waiter) startWorking() {
	w.atWork = 1
	for w.atWork == 1 {
		didATask := false

		//Look for table that needs their order taken
		order := diningHall.tableList.serveTable(w)

		for success := false; order != nil && !success; {
			//Send order
			success = diningHall.sendOrder(order)
			if success {
				didATask = true
				w.modifierId = order.Id
				w.statusId = 1
				time.Sleep(timeUnit)
			} else {
				w.statusId = 3
				time.Sleep(timeUnit)
			}
		}

		//Receive delivery
		select {
		case delivery := <-diningHall.deliveryChan:
			didATask = true
			//Serve delivery to the required table
			w.statusId = 2
			w.modifierId = delivery.OrderId
			time.Sleep(timeUnit)
			now := getUnixTimeUnits()
			diningHall.tableList.deliver(delivery, now)
		default:
			break
		}

		if !didATask {
			//Wait one unit because there are no tasks
			w.statusId = 0
			time.Sleep(timeUnit)
		}

	}
}
func (w *Waiter) getStatus() string {
	status := "Waiter id:" + strconv.Itoa(w.id) + " Status:" + waiterStatus[w.statusId]
	if w.statusId == 1 || w.statusId == 2 {
		return status + strconv.Itoa(w.modifierId)
	}
	return status
}
