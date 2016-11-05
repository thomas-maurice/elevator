package algorithms

import (
	"github.com/Sirupsen/logrus"
	"github.com/thomas-maurice/elevator/scheduler"
)

// Sooner Available Append Scheduler
// The goal of this algo is to assign the pickup and drop off of the passenger
// To the elevator that will be free first AFTER serving the request, regardless of
// how many requests are piled up
func SoonerAvailableAppendScheduler(self *scheduler.Scheduler) {
	if len(self.Requests) == 0 {
		return
	}
	logrus.Info("Scheduling using SAAS algorithm")
	for _, request := range self.Requests {
		logrus.WithFields(logrus.Fields{"requested_floor": request}).Info("Processing request")
		bestCandidate := self.Elevators[0]
		for _, candidate := range self.Elevators {
			candidateSteps, _ := candidate.StepsUntilIdleWithAppendedRequest(request)
			bestCandidateSteps, _ := bestCandidate.StepsUntilIdleWithAppendedRequest(request)
			if candidateSteps < bestCandidateSteps {
				bestCandidate = candidate
			}
			if bestCandidateSteps == 0 {
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
