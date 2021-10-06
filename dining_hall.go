package main

import "time"

type DiningHall struct {
	diningHallWeb DiningHallWeb
	waiterList    WaiterList
	tableList     TableList
	connected     bool
}

func (dh *DiningHall) start() {
	go dh.tryConnectKitchen()
	dh.diningHallWeb.start()
}

func (dh *DiningHall) connectionSuccessful() {
	dh.connected = true
	dh.tableList.start()
	dh.waiterList.start()
}

func (dh *DiningHall) tryConnectKitchen() {
	dh.connected = false
	for dh.connected {
		if dh.diningHallWeb.establishConnection() {
			dh.connectionSuccessful()
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (dh *DiningHall) sendOrder(order *Order) bool {
	return dh.diningHallWeb.sendOrder(order)
}
