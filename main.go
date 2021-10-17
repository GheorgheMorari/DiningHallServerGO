package main

import (
	"os"
	"time"
)

var kitchenServerHost = "http://localhost"

const diningHallPort = ":7500"
const kitchenServerPort = ":8000"

const tableN = 6
const waiterN = 3

const timeUnit = time.Millisecond

var diningHall DiningHall

func main() {

	args := os.Args
	if len(args) > 1 {
		//Set the docker internal host
		kitchenServerHost = args[1]
	}

	diningHall.start()
}
