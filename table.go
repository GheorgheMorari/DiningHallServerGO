package main

import (
	"math/rand"
	"sync/atomic"
	"time"
	"unsafe"
)

type Table struct {
	id        int
	ordered   int32
	occupied  int32
	available int32
	order     *Order
}

func (t Table) startAvailability() {
	atomic.StoreInt32(&t.available, 1)
}

func (t Table) deliver(delivery *Delivery) {
	//Wait based on the delivery size
	time.Sleep(time.Second * time.Duration(len(delivery.Items)))

	//TODO Add reputation calculation system

	atomic.StoreInt32(&t.ordered, 0)
	atomic.StoreInt32(&t.occupied, 0)
	t.waitCustomers()
}

func (t *Table) waitCustomers() {
	if t.available == 1 && t.occupied == 0 {
		atomic.StoreInt32(&t.ordered, 1)
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(t.order)), unsafe.Pointer(generateOrder(t)))
		atomic.StoreInt32(&t.ordered, 0)
		atomic.StoreInt32(&t.occupied, 1)
	}
}

func (t Table) stopAvailability() {
	atomic.StoreInt32(&t.available, 0)
}

func (t Table) waitForOrderList(){
	atomic.StoreInt32(&t.ordered, 1)
	time.Sleep(time.Second * 2) //Wait 2 seconds for the order list to free
	atomic.StoreInt32(&t.ordered, 0)
}
