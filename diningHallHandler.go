package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DiningHallHandler struct {
	packetsReceived int32
	postReceived    int32
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
				fmt.Fprint(w, response)
				return
			} else {
				latestDelivery = new(Delivery)
				var requestBody = make([]byte, r.ContentLength)
				r.Body.Read(requestBody)
				json.Unmarshal(requestBody, latestDelivery)
				diningHall.diningHallWeb.setDelivery(latestDelivery)

				//Respond with "OK"
				fmt.Fprint(w, response)
			}
		}
	case http.MethodGet:
		{
			//Todo display what dining hall is doing

			fmt.Fprintln(w, "DiningHall server is UP on port "+diningHallPort)
			if diningHall.connected {
				fmt.Fprintln(w, "DiningHall successfully connected to kitchen on address:"+kitchenServerHost+kitchenServerPort)
			} else {
				fmt.Fprintln(w, "DiningHall did not establish connection to kitchen on address:"+kitchenServerHost+kitchenServerPort)
			}
			if r.RequestURI == "/connect" {
				diningHall.diningHallWeb.establishConnection()
			}
		}
	case http.MethodConnect:
		{
			diningHall.connectionSuccessful()
			fmt.Fprint(w, "OK")
		}
	default:
		{
			fmt.Fprintln(w, "UNSUPPORTED METHOD")
		}
	}
}
