package algorithms

import (
	"github.com/Sirupsen/logrus"
	"github.com/thomas-maurice/elevator/scheduler"
)

// Sooner Available Insert ASAP Scheduler
// The goal of this algo is to assign the pickup and drop off of the passenger
// To the elevator that will be free first AFTER serving the request, regardless of
// how many requests are piled up. Note thatr this time the request is inserted inside the scheduler's queue.
func SoonerAvailableInsertASAPScheduler(self *scheduler.Scheduler) {
	if len(self.Requests) == 0 {
		return
	}
	logrus.Info("Scheduling using SAISAPS algorithm")
	for _, request := range self.Requests {
		logrus.WithFields(logrus.Fields{"requested_floor": request}).Info("Processing request")
		bestCandidate := self.Elevators[0]
		evaluationElevator := scheduler.Elevator{
			NextStops: scheduler.InsertASAP(bestCandidate.NextStops, request),
			Floor:     bestCandidate.Floor,
		}
		bestCandidateSteps, _ := evaluationElevator.StepsUntilIdle()
		for _, candidate := range self.Elevators {
			evaluationElevator := scheduler.Elevator{
				NextStops: scheduler.InsertASAP(candidate.NextStops, request),
				Floor:     candidate.Floor,
			}
			candidateSteps, _ := evaluationElevator.StepsUntilIdle()
			if candidateSteps < bestCandidateSteps {
				bestCandidate = candidate
			}
			if bestCandidateSteps == 0 {
				// No need for further processing
				break
			}
		}
		logrus.WithFields(logrus.Fields{"requested_floor": request}).Infof("Elected elevator %d", bestCandidate.Id)
		bestCandidate.InsertRequest(request)
	}
	self.Requests = []scheduler.PickUpRequest{}
}
