package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
)

type DiningHallHandler struct{}

func (DiningHallHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//TODO make buffer static
		var buffer = make([]byte, r.ContentLength)
		r.Body.Read(buffer)
		fmt.Fprintln(w, "Dining hall http request post method detected")
		fmt.Fprintln(w, "Post Method Body:\n"+string(buffer))
	} else {
		if r.Method == http.MethodGet {
			if r.RequestURI == "/start" {
				startFakeOrders(w, r)
			}
			if r.RequestURI == "/send" {
				sendOneFakeOrder(w, r)
			}
			if r.RequestURI == "/stop" {
				stopFakeOrders(w, r)
			}
			fmt.Fprintln(w, "Dining Hall server is UP on port "+diningHallPort)
		}
	}

}

const diningHallPort = ":8080"
const kitchenServerPort = ":8000"

var runFakeOrders int32 = 0

func main() {
	rand.Seed(69)

	//TODO send Connect request ensure connection

	var diningHallServer http.Server
	diningHallServer.Addr = diningHallPort
	diningHallServer.Handler = DiningHallHandler{}

	fmt.Println("DiningHallServer is listening and serving")
	if err := diningHallServer.ListenAndServe(); err != nil {
		//Stop sending fake orders
		atomic.StoreInt32(&runFakeOrders, 0)
		log.Fatal(err)
	}
}
