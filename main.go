package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/thomas-maurice/elevator/algorithms"
	"github.com/thomas-maurice/elevator/scheduler"
	"time"
)

func main() {
	logrus.Info("Ready to transport people!")

	sched := scheduler.NewScheduler()
	go sched.Router.Run(":8080")
	sched.AddElevator(scheduler.NewElevator(1))
	sched.AddElevator(scheduler.NewElevator(2))

	/*sched.AddRequest(scheduler.PickUpRequest{Source: 0, Destination: 10})
	sched.AddRequest(scheduler.PickUpRequest{Source: 10, Destination: 8})
	sched.AddRequest(scheduler.PickUpRequest{Source: 4, Destination: 2})
	sched.AddRequest(scheduler.PickUpRequest{Source: 2, Destination: 4})*/

	for {
		sched.Schedule(algorithms.SoonerAvailableInsertASAPScheduler)
		time.Sleep(time.Second * 1)
	}
}
