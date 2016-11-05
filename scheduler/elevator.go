package scheduler

import (
	"github.com/Sirupsen/logrus"
	"math"
)

const (
	DirectionUp = iota
	DirectionDown
	DirectionIdle
)

type Elevator struct {
	Id        int64   `json:"id"`         // Elevator ID
	Floor     int64   `json:"floor"`      // What floor is the elevator in
	NextStops []int64 `json:"next_stops"` // List of floors the elevator has to get to
}

// Returns a new Elevator initialized with an Id
func NewElevator(id int64) *Elevator {
	return &Elevator{
		Id: id,
	}
}

// Removes elements that are sonsecutive and the same
func (self *Elevator) OptimizeStops() {
	self.NextStops = OptimizePath(self.NextStops)
}

// Returns te direction in which the elevator is going
func (self *Elevator) GetDirection() int64 {
	if len(self.NextStops) == 0 {
		return DirectionIdle
	} else if self.NextStops[0] < self.Floor {
		return DirectionDown
	}

	return DirectionUp
}

// Insert a request inside the queue
func (self *Elevator) InsertRequest(request PickUpRequest) {
	self.NextStops = InsertASAP(self.NextStops, request)
}

// Returns a logger for debugging purposes
func (self *Elevator) GetLogger() *logrus.Entry {
	untilIdle, _ := self.StepsUntilIdle()
	return logrus.WithFields(
		logrus.Fields{
			"elevator_id":   self.Id,
			"current_floor": self.Floor,
			"idle_in":       untilIdle,
		})
}

// Appends a pickup and a destination to the list
func (self *Elevator) AppendStop(stop int64) {
	self.NextStops = append(self.NextStops, stop)
	self.OptimizeStops()
}

// Returns how many steps it will take until the elevator is idle
// and the final floor it will be
func (self *Elevator) StepsUntilIdle() (int64, int64) {
	return ComputePathCost(self.Floor, self.NextStops)
}

// Returns how many steps it will take until the elevator is idle and the final
// floor it will be if we pickup the request and append it to the queue
func (self *Elevator) StepsUntilIdleWithAppendedRequest(request PickUpRequest) (int64, int64) {
	steps, currentFloor := self.StepsUntilIdle()

	if currentFloor != request.Source {
		steps += 1
	}
	// +1 because 1 step to let the people flow
	steps += int64(math.Abs(float64(request.Destination + 1 - currentFloor)))

	return steps, request.Destination
}

// Performs a step.
func (self *Elevator) Step() {
	logger := self.GetLogger()
	// If the elevator is not working, then noop
	if len(self.NextStops) == 0 {
		logger.Info("I am idle")
		return
	}

	// Have we reached our next goal ?
	if self.NextStops[0] == self.Floor {
		// Update the NextStops, to remove the curent one
		// and return to emulate a one step stop
		self.NextStops = self.NextStops[1:]
		logger.Infof("People flow for one step")
		return
	}

	switch self.GetDirection() {
	case DirectionUp:
		logger.WithFields(logrus.Fields{"target": self.NextStops[0]}).Infof("Going up")
		self.Floor++
	case DirectionDown:
		logger.WithFields(logrus.Fields{"target": self.NextStops[0]}).Infof("Going down")
		self.Floor--
	case DirectionIdle:
		logger.Infof("Doing nothing")
		// noop
	}

	return
}
