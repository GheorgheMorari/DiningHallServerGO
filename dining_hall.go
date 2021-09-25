package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func sendOneFakeOrder(w http.ResponseWriter, r *http.Request) {
	var diningHallClient http.Client
	fmt.Fprintln(w, "Sent one fake order")

	var requestBody = []byte(testingPayload)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost"+kitchenServerPort, bytes.NewBuffer(requestBody))
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
	fmt.Fprintln(w, "Started sending fake orders")
	go sendFakeOrders(&runFakeOrders)
}
func stopFakeOrders(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&runFakeOrders, 0)
	fmt.Fprintln(w, "Stopped sending fake orders")
}

func sendFakeOrders(runFakeOrders *int32) {
	errorCount := 0
	requestCount := 0
	var diningHallClient http.Client
	var requestBody = []byte(testingPayload)
	for *runFakeOrders == 1 {
		//TODO handle errors and requests
		request, _ := http.NewRequest(http.MethodPost, "http://localhost"+kitchenServerPort, bytes.NewBuffer(requestBody))

		_, err := diningHallClient.Do(request)

		if err != nil {
			errorCount++
			fmt.Println("Error, encountered: ", err)
			fmt.Println("Requests:", requestCount, " Errors:", errorCount)
		} else {
			requestCount++
			if requestCount%5 == 0 {
				fmt.Println("Requests:", requestCount, " Errors:", errorCount)
			}
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

const testingPayload = `{"order_id": 1,
"table_id": 1,
"waiter_id": 1,
"items": [ 3, 4, 4, 2 ],
"priority": 3,
"max_wait": 45,
"pick_up_time": 1631453140 
}`
