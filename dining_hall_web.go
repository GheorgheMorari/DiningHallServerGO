package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type DiningHallWeb struct {
	diningHallServer  http.Server
	diningHallHandler DiningHallHandler
	diningHallClient  http.Client
	latestDelivery    *Delivery
	deliveryMutex     sync.Mutex
}

func (dhw *DiningHallWeb) getDelivery() *Delivery {
	dhw.deliveryMutex.Lock()
	defer dhw.deliveryMutex.Unlock()

	ret := dhw.latestDelivery
	dhw.latestDelivery = nil

	return ret
}

func (dhw *DiningHallWeb) setDelivery(delivery *Delivery) {
	dhw.deliveryMutex.Lock()
	defer dhw.deliveryMutex.Unlock()

	dhw.latestDelivery = delivery

}

func (dhw *DiningHallWeb) start() {
	dhw.diningHallServer.Addr = diningHallPort
	dhw.diningHallServer.Handler = &dhw.diningHallHandler

	fmt.Println(time.Now())
	fmt.Println("DiningHallServer is listening and serving on port:" + diningHallPort)
	if err := dhw.diningHallServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (dhw *DiningHallWeb) sendOrder(order *Order) bool {
	requestBody := order.getPayload()
	request, _ := http.NewRequest(http.MethodPost, kitchenServerHost+kitchenServerPort+"/order", bytes.NewBuffer(requestBody))
	response, err := dhw.diningHallClient.Do(request)

	if err != nil {
		fmt.Println("Could not send order to kitchen.")
		log.Fatal(err)
		return false
	}
	var responseBody = make([]byte, response.ContentLength)
	response.Body.Read(responseBody)
	if string(responseBody) != "OK" {
		return false
	}

	return true
}

func (dhw *DiningHallWeb) establishConnection() bool{
	if diningHall.connected == true {
		return false
	}
	request, _ := http.NewRequest(http.MethodConnect, kitchenServerHost+kitchenServerPort+"/", bytes.NewBuffer([]byte{}))
	response, err := dhw.diningHallClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var responseBody = make([]byte, response.ContentLength)
	response.Body.Read(responseBody)
	if string(responseBody) != "OK" {
		return false
	}

	return true
}

//
//var diningHallClient http.Client
//func sendOneFakeOrder(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Sent one fake order")
//
//	order := generateOrder()
//	var requestBody = order.getPayload()
//	request, _ := http.NewRequest(http.MethodPost, kitchenServerHost+kitchenServerPort+"/order", bytes.NewBuffer(requestBody))
//	response, err := diningHallClient.Do(request)
//
//	if err != nil {
//		fmt.Fprintln(w, "ERROR DETECTED:", err)
//	} else {
//		fmt.Fprintln(w, "Response detected.")
//		var buffer = make([]byte, response.ContentLength)
//		response.Body.Read(buffer)
//		fmt.Fprintln(w, "Response Body:\n"+string(buffer))
//	}
//}
//func startFakeOrders(w http.ResponseWriter, r *http.Request) {
//	atomic.StoreInt32(&runFakeOrders, 1)
//	threads := 5
//	diningHallClient.CloseIdleConnections()
//	fmt.Fprintf(w, "Started sending fake orders on %d threads.\n", threads)
//	for i := 0; i < threads; i++ {
//		go sendFakeOrders(&runFakeOrders)
//	}
//}
//
//func stopFakeOrders(w http.ResponseWriter, r *http.Request) {
//	atomic.StoreInt32(&runFakeOrders, 0)
//	fmt.Fprintln(w, "Stopped sending fake orders.")
//	diningHallClient.CloseIdleConnections()
//}
//
//func sendFakeOrders(runFakeOrders *int32) {
//	//var diningHallClient http.Client
//	for *runFakeOrders == int32(1) {
//
//		order := generateOrder()
//		var requestBody = order.getPayload()
//		request, _ := http.NewRequest(http.MethodPost, kitchenServerHost+kitchenServerPort+"/order", bytes.NewBuffer(requestBody))
//		_, err := diningHallClient.Do(request)
//
//		if err != nil {
//			fmt.Println("Thread finished sending messages, due to error:")
//			fmt.Println(err)
//			return
//		}
//		time.Sleep(time.Duration(rand.Float32()*3+1)*time.Second)
//	}
//	fmt.Println("Thread finished sending messages, the sending of the requests was stopped manually.")
//}
