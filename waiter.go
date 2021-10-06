package main

import "time"

type Waiter struct {
	id int
	atWork int
}

func (w Waiter) startWorking() {
	w.atWork = 1
	for w.atWork == 1{
		didATask := false

		//Look for table that needs their order taken
		table := diningHall.tableList.lookUpTable()

		if table != nil {
			//Get order
			order := table.order
			order.WaiterId = w.id

			//Send order
			success := diningHall.sendOrder(order)
			didATask = true
			if !success{
				go table.waitForOrderList()
				didATask = false
			}
		}

		//Receive delivery
		delivery := diningHall.diningHallWeb.getDelivery()
		if delivery != nil {
			didATask = true
			//Serve delivery to the required table
			go diningHall.tableList.tableList[delivery.TableId-1].deliver(delivery)
		}


		if !didATask {
			//Wait one second because there are no tasks
			time.Sleep(time.Second)
		}

	}
}

