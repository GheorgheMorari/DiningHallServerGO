package main

import (
	"fmt"
	"time"
)

type DiningHall struct {
	diningHallWeb DiningHallWeb
	waiterList    *WaiterList
	tableList     *TableList
	deliveryChan  chan *Delivery
	ratings       *Rating
	connected     bool
	startTime     time.Time
}

func (dh *DiningHall) start() {
	dh.ratings = NewRating()
	dh.waiterList = NewWaiterList()
	dh.tableList = NewTableList()
	go dh.tryConnectKitchen()
	dh.diningHallWeb.start()
}

func (dh *DiningHall) connectionSuccessful() {
	dh.startTime = time.Now()
	if dh.connected {
		return
	}
	dh.connected = true
	dh.deliveryChan = make(chan *Delivery)
	dh.tableList.start()
	dh.waiterList.start()
}

func (dh *DiningHall) tryConnectKitchen() {
	dh.connected = false
	for !dh.connected {
		if dh.diningHallWeb.establishConnection() {
			dh.connectionSuccessful()
			break
		} else {
			time.Sleep(timeUnit)
		}
	}
}

func (dh *DiningHall) sendOrder(order *Order) bool {
	return dh.diningHallWeb.sendOrder(order)
}

func (dh *DiningHall) getStatus() string {
	ret := "Running for:" + fmt.Sprintf("%v", time.Since(dh.startTime))
	ret += makeDiv("Rating:" + fmt.Sprintf("%f", dh.ratings.getAverage()) + " Total reviews:" + fmt.Sprintf("%d", dh.ratings.getNumOfOrders()))
	ret += "Waiters:"
	for _, waiter := range dh.waiterList.waiterList {
		ret += makeDiv(waiter.getStatus())
	}
	ret += "Tables:"
	for _, table := range dh.tableList.tableList {
		ret += makeDiv(table.getStatus())
	}

	return ret
}
