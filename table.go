package main

import (
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var tableStatus = [...]string{"Waiting for delivery.", "Waiting for customers.", "Eating.", "Waiting for kitchen's order list to empty.", "Waiting for waiter."}

type Table struct {
	id        int
	ordered   int32
	occupied  int32
	available int32
	statusId  int
	order     *Order
}

func NewTable(id int, ordered int32, occupied int32, available int32, statusId int, order *Order) *Table {
	ret := new(Table)
	ret.id = id
	ret.ordered = ordered
	ret.occupied = occupied
	ret.available = available
	ret.statusId = statusId
	ret.order = order
	return ret
}

func (t *Table) startAvailability() {
	atomic.StoreInt32(&t.available, 1)
	t.waitCustomers()
}

func (t *Table) deliver(delivery *Delivery, now int64) {
	//Wait based on the delivery size
	t.statusId = 2
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
	time.Sleep(timeUnit * (time.Duration(len(delivery.Items) + 1)))

	atomic.StoreInt32(&t.ordered, 0)
	atomic.StoreInt32(&t.occupied, 0)
	t.waitCustomers()
}

func (t *Table) waitCustomers() {
	if t.available == 1 && t.occupied == 0 {
		syncMutex := sync.Mutex{}
		atomic.StoreInt32(&t.ordered, 1)
		t.statusId = 1
		time.Sleep(timeUnit * time.Duration(rand.Intn(10)))

		syncMutex.Lock()
		if t.order == nil {
			t.order = generateOrder(t)
		}
		syncMutex.Unlock()
		//addr := (*unsafe.Pointer)(unsafe.Pointer(t.order))
		//newOrder := unsafe.Pointer(generateOrder(t))
		//atomic.StorePointer(addr, newOrder)
		atomic.StoreInt32(&t.ordered, 0)
		atomic.StoreInt32(&t.occupied, 1)
		t.statusId = 4
	}
}

func (t *Table) getOrder(waiter *Waiter) *Order {
	t.statusId = 0
	t.order.WaiterId = waiter.id
	return t.order
}

func (t *Table) stopAvailability() {
	atomic.StoreInt32(&t.available, 0)
}

func (t *Table) waitForOrderList() {
	atomic.StoreInt32(&t.ordered, 1)
	t.statusId = 3
	time.Sleep(timeUnit * 2) //Wait 2 units for the order list to free
	atomic.StoreInt32(&t.ordered, 0)
}

func (t *Table) getStatus() string {
	waitStatus := ""
	if t.occupied == 1 && t.ordered == 1 && t.order != nil {
		waitStatus = " Waiting for:" + strconv.Itoa(int(getUnixTimeUnits()-t.order.PickUpTime)) + "sec" + " Max wait:" + strconv.Itoa(t.order.MaxWait)
	}
	return "Table id:" + strconv.Itoa(t.id) + " Status:" + tableStatus[t.statusId] + waitStatus
}
