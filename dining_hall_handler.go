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
			latestDelivery := new(Delivery)
			var requestBody = make([]byte, r.ContentLength)
			r.Body.Read(requestBody)
			json.Unmarshal(requestBody, latestDelivery)
			diningHall.deliveryChan <- latestDelivery

			//Respond with "OK"
			fmt.Fprint(w, "OK")
		}
	case http.MethodGet:
		{
			fmt.Fprintln(w, "<head><meta http-equiv=\"refresh\" content=\"1\" /></head>")
			fmt.Fprintln(w, makeDiv("DiningHall server is UP on port "+diningHallPort))
			if diningHall.connected {
				fmt.Fprintln(w, makeDiv("DiningHall successfully connected to kitchen on address:"+kitchenServerHost+kitchenServerPort))
			} else {
				fmt.Fprintln(w, makeDiv("DiningHall did not establish connection to kitchen on address:"+kitchenServerHost+kitchenServerPort))
				err := diningHall.diningHallWeb.connectionError
				if err != nil {
					fmt.Fprintln(w, makeDiv("Connection error: "+err.Error()))
				}
			}
			fmt.Fprintln(w, makeDiv(diningHall.getStatus()))
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
