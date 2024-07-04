package main

import (
	"fmt"
	"generic-patterns/behavioral"
)

type StationManager struct {
	behavioral.Mediator
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	behavioral.Component[StationManager]
}

func (g *PassengerTrain) arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

type FreightTrain struct {
	behavioral.Component[StationManager]
}

func (g *FreightTrain) arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

func MainMediatorExample() {

	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{}
	passengerTrain.Register(stationManager)

	freightTrain := &FreightTrain{}
	freightTrain.Register(stationManager)

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()

}
