package main

import (
	"fmt"
	"generic-patterns/creational"
)

type Transport interface {
	deliver()
}

type Truck struct{}

func (t *Truck) deliver() {
	println("Delivering by truck")
}

type Ship struct{}

func (s *Ship) deliver() {
	println("Delivering by ship")
}

func GenerateTransportFactory(transportType string) (*creational.Factory[Transport], error) {

	// make function for the factory
	transportMakeFunc := func() Transport {
		switch transportType {
		case "Truck":
			return &Truck{}
		case "Ship":
			return &Ship{}
		default:
			return nil
		}
	}

	// Generate a factory for the given transport type
	factory := creational.NewFactory(transportMakeFunc)

	if factory != nil {
		return factory, nil
	}

	return nil, fmt.Errorf("UNKNOWN TRANSPORT TYPE")
}

func MainFactoryExample() {
	// Generate a factory for trucks
	TruckFactory, _ := GenerateTransportFactory("Truck")

	// Make a truck
	truck1 := TruckFactory.Make()
	truck2 := TruckFactory.Make()

	// Deliver by truck
	truck1.deliver()
	truck2.deliver()
}
