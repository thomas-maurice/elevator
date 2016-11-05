package scheduler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

type Scheduler struct {
	Elevators []*Elevator     `json:"elevators"`
	Requests  []PickUpRequest `json:"requests"`
	Lock      *sync.Mutex
	Router    *gin.Engine
}

type PickUpRequest struct {
	Source      int64 `json:"source"`      // Source floor
	Destination int64 `json:"destination"` // Destination floor
}

func NewScheduler() *Scheduler {
	sched := &Scheduler{
		Lock: new(sync.Mutex),
	}

	router := gin.Default()

	router.POST("/pickup/:start/:end", func(c *gin.Context) {
		start, err := strconv.ParseInt(c.Param("start"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid start floor"})
			return
		}

		end, err := strconv.ParseInt(c.Param("end"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid end floor"})
			return
		}

		sched.AddRequest(PickUpRequest{Source: start, Destination: end})

		c.JSON(200, gin.H{"message": "yeah"})
	})

	sched.Router = router
	return sched
}

func (self *Scheduler) AddElevator(e *Elevator) {
	self.Elevators = append(self.Elevators, e)
}

func (self *Scheduler) AddRequest(request PickUpRequest) {
	self.Lock.Lock()
	defer self.Lock.Unlock()
	self.Requests = append(self.Requests, request)
}

func (self *Scheduler) Schedule(fn func(*Scheduler)) {
	self.Lock.Lock()
	defer self.Lock.Unlock()

	fn(self)
	self.Step()
}

func (self *Scheduler) Step() {
	for _, elevator := range self.Elevators {
		//logrus.WithFields(logrus.Fields{"elevator_id": elevator.Id}).Info("Performing step on elevator")
		elevator.Step()
	}
}
