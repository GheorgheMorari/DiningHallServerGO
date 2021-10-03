package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

var diningHallClient http.Client
func sendOneFakeOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sent one fake order")

	order := getRandomOrder()
	var requestBody = order.getPayload()
	request, _ := http.NewRequest(http.MethodPost, kitchenServerHost+kitchenServerPort+"/order", bytes.NewBuffer(requestBody))
	response, err := diningHallClient.Do(request)

	if err != nil {
		fmt.Fprintln(w, "ERROR DETECTED:", err)
	} else {
		fmt.Fprintln(w, "Response detected.")
		var buffer = make([]byte, response.ContentLength)
		response.Body.Read(buffer)
		fmt.Fprintln(w, "Response Body:\n"+string(buffer))
	}
}
func startFakeOrders(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&runFakeOrders, 1)
	threads := 5
	diningHallClient.CloseIdleConnections()
	fmt.Fprintf(w, "Started sending fake orders on %d threads.\n", threads)
	for i := 0; i < threads; i++ {
		go sendFakeOrders(&runFakeOrders)
	}
}

func stopFakeOrders(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&runFakeOrders, 0)
	fmt.Fprintln(w, "Stopped sending fake orders.")
	diningHallClient.CloseIdleConnections()
}

func sendFakeOrders(runFakeOrders *int32) {
	//var diningHallClient http.Client
	for *runFakeOrders == int32(1) {

		order := getRandomOrder()
		var requestBody = order.getPayload()
		request, _ := http.NewRequest(http.MethodPost, "http://localhost"+kitchenServerPort+"/order", bytes.NewBuffer(requestBody))
		_, err := diningHallClient.Do(request)

		if err != nil {
			fmt.Println("Thread finished sending messages, due to error:")
			fmt.Println(err)
			return
		}
		time.Sleep(time.Duration(rand.Float32()*3+1)*time.Second)
	}
	fmt.Println("Thread finished sending messages, the sending of the requests was stopped manually.")
}

const testingPayload = `{"order_id": 1,
"table_id": 1,
"waiter_id": 1,
"items": [ 3, 4, 4, 2 ],
"priority": 3,
"max_wait": 45,
"pick_up_time": 1631453140 
}`
