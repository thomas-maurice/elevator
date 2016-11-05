# Elevators

Simple elevator scheduling framework

## Intro
It has been built in 5 steps.

* Definition of the data structures
* Implementation of the first come first served algo
* Implementation of the Sooner available scheduler algo
* Implementation of the Sooner available append scheduler algo
* Implementation of the Sooner available insert ASAP scheduler algo

This simulator works in steps. At each steps the scheduler will take waiting requests
and feed them to the scheduling algo (editable in main.go).

The scheduler can operate as many elevators as you want.

Then each second, it will perform a step on the elevators. An elevator can do the following things:

* Go up
* Go down
* Stay there to let people flow
* nothing

Each of these action costs one step for simplicity's sake.

## First come, first served
This algorithm is über simple. Each time a new request arrives. The scheduler gives it
to the elevator that has the smallest waiting queue, regardless of the cost of the steps

The next algorithm have been built with the following paradigm: operating elevators is expensive, so
I want to minimize the time they will be moving. For that, at a given point in time I want to schedule
them so that they will individually run for the shortest period of time as possible. This may not
be optimal for people in these elevators, but ut is for me as an elevator provider <:o)

Also requests are processed sequentially, the request n+1 will NOT have any effect on the scheduling
of request n.

## Sooner available scheduler
This one is simple too. We want to assign the incomming request to the elevator that will become idle
the first.

## Sooner available append scheduler
This is an adaptation of the previous one, in which we do not assign the request to the elevator
that will be come available first, but to the elevator that will become available first **if it is
assigned the request**, which is a bit more efficient.

## Sooner available insert ASAP scheduler
This one is an enhancement of the previous algorithm. It will try to insert the requested pickup inside
the elevator's path. Meaning that for instance if it is at floor 0, and has to go to floor 5, but someone
wanting to go from floor 2 to 3, then the elevator's path will be altered as follows: 0 -> 2 -> 3 -> 5.

# Some benchmarks

We have the following setup: 2 elevators and the following requests:

* 0 -> 10
* 10 -> 2
* 4 -> 2
* 2 -> 4

The results of the benchmark are the "available_in" values of the elevators right after scheduling.

* FCFS: e1: 22 e2: 24
* SAS: e1: 22 e2: 24
* SAAS: e1: 22 e2: 24
* SAIASAPS: e1: 19 e2: 14

# With your own datasets

You can programatically feed input to the scheduler, for that:
```
curl -XPOST localhost:8080/pickup/<start>/<end>
```

The logging is explicit enough for you to track what's going on
