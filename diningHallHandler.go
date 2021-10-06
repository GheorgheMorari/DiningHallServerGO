package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DiningHallHandler struct {
	packetsReceived      int32
	postReceived         int32
}

func (d DiningHallHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		{
			response := "OK"
			latestDelivery := diningHall.diningHallWeb.getDelivery()
			if latestDelivery != nil {
				diningHall.diningHallWeb.setDelivery(latestDelivery)
				response = "NOT OK"
				fmt.Fprintln(w, response)
				return
			} else {
				latestDelivery = new(Delivery)
				var requestBody = make([]byte, r.ContentLength)
				r.Body.Read(requestBody)
				json.Unmarshal(requestBody, latestDelivery)
				diningHall.diningHallWeb.setDelivery(latestDelivery)

				//Respond with "OK"
				fmt.Fprintln(w, response)
			}
		}
	case http.MethodGet:
		{
			//Todo display what dining hall is doing

			//fmt.Fprintln(w, "Dining Hall server is UP on port "+diningHallPort)
			//if r.RequestURI == "/start" {
			//	startFakeOrders(w, r)
			//}
			//if r.RequestURI == "/send" {
			//	sendOneFakeOrder(w, r)
			//}
			//if r.RequestURI == "/stop" {
			//	stopFakeOrders(w, r)
			//}
		}
	case http.MethodConnect:
		{
			diningHall.diningHallWeb.establishConnection()
			fmt.Fprintln(w, "OK")
		}
	default:
		{
			fmt.Fprintln(w, "UNSUPPORTED METHOD")
		}
	}
}

