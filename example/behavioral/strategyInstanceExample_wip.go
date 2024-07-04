package main

import "github.com/Zando74/generic-patterns/behavioral"

type Trip struct {
}

func (t *Trip) Travel(strategy behavioral.StrategyInstance[Trip]) {
	strategy.Execute()
}

type Plane struct {
	// Complex structure...
}

type FlightStrategy struct {
	Plane *Plane
	behavioral.StrategyInstance[Trip]
}

func (f *FlightStrategy) Execute() {
	println("Setup the plane")
	println("Do some checking")
	println("Traveling by flight")
}

func MainStrategyInstanceExample() {

	trip := &Trip{}

	plane := &Plane{}
	byFlight := &FlightStrategy{
		Plane: plane,
	}

	trip.Travel(byFlight)
}
