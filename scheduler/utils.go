package scheduler

import (
	"math"
)

// Returns the optimized path by removing two consecutive identical elements
func OptimizePath(path []int64) []int64 {
	newStops := []int64{}
	prev := int64(0)
	for idx, stop := range path {
		if idx == 0 || prev != stop {
			newStops = append(newStops, stop)
			prev = stop
		}
	}
	return newStops
}

// Computes the cost of a path in steps, from a starting point to a destination
func ComputePathCost(start int64, path []int64) (int64, int64) {
	steps := 0.0
	currentFloor := start
	for _, stop := range path {
		steps++ // Because we have one step wait when we reach a destination
		steps += math.Abs(float64(stop - currentFloor))
		currentFloor = stop
	}
	return int64(steps), currentFloor
}

func InsertASAP(path []int64, request PickUpRequest) []int64 {
	newPath := []int64{}
	addedStart := false
	addedStop := false
	prev := int64(0)
	for idx, elt := range path {
		if idx == 0 {
			newPath = append(newPath, elt)
			prev = elt
			continue
		}
		if !addedStart {
			// If we are going down
			if prev <= elt && request.Source <= elt {
				newPath = append(newPath, request.Source)
				addedStart = true
				// If the destination is too in the interval, let's take advantage of it
				if request.Destination <= elt {
					newPath = append(newPath, request.Destination)
					addedStop = true
				}
			} else if prev >= elt && request.Source >= elt { // We are going up
				newPath = append(newPath, request.Source)
				addedStart = true
				// If the destination is too in the interval, let's take advantage of it
				if request.Destination >= elt {
					newPath = append(newPath, request.Destination)
					addedStop = true
				}
			}
		} else if addedStart && !addedStop {
			// We only have the stopping point to add
			// If we are going down
			if prev <= elt && request.Destination <= elt {
				newPath = append(newPath, request.Destination)
				addedStop = true
			} else if prev >= elt && request.Destination >= elt { // We are going up
				newPath = append(newPath, request.Destination)
				addedStop = true
			}
		}
		newPath = append(newPath, elt)
	}

	if !addedStart {
		newPath = append(newPath, request.Source)
	}
	if !addedStop {
		newPath = append(newPath, request.Destination)
	}

	return OptimizePath(newPath)
}
