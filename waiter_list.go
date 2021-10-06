package main

type WaiterList struct {
	waiterList      []Waiter
	waiterIdCounter int
}

func (wl WaiterList) start() {
	wl.waiterIdCounter = 0
	wl.waiterList = []Waiter{}
	for i := 0; i < waiterN; i++ {
		wl.waiterList = append(wl.waiterList, Waiter{wl.waiterIdCounter,0})
		wl.waiterIdCounter++
	}

	for _, waiter := range wl.waiterList {
		go waiter.startWorking()
	}
}
