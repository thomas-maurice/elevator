package algorithms

import (
	"github.com/Sirupsen/logrus"
	"github.com/thomas-maurice/elevator/scheduler"
)

// First come served algorithm
// Here we just pile up the requests on the elevator that has the least
// ammount of requests stacked.
func FCFSScheduler(self *scheduler.Scheduler) {
	if len(self.Requests) == 0 {
		return
	}
	logrus.Info("Scheduling using FCFS algorithm")
	for _, request := range self.Requests {
		logrus.WithFields(logrus.Fields{"requested_floor": request}).Info("Processing request")
		bestCandidate := self.Elevators[0]
		for _, candidate := range self.Elevators {
			if len(candidate.NextStops) < len(bestCandidate.NextStops) {
				bestCandidate = candidate
			}
			if len(bestCandidate.NextStops) == 0 {
				// No need for further processing
				break
			}
		}
		logrus.WithFields(logrus.Fields{"requested_floor": request}).Infof("Elected elevator %d", bestCandidate.Id)
		bestCandidate.AppendStop(request.Source)
		bestCandidate.AppendStop(request.Destination)
	}
	self.Requests = []scheduler.PickUpRequest{}
}
