package main

import (
	"fmt"
	"generic-patterns/creational"
)

// DBConnection is a struct that represents a database connection
// It should be a singleton
type DBConnection struct {
	Name string
	// others attributes ...
}

// NewDBConnection creates a new instance of DBConnection
func NewDBConnection() *DBConnection {
	return &DBConnection{Name: "Unique instance"}
}

// creation.NewSingleton creates a new singleton instance
var DBConnectionSingleton = creational.NewSingleton(NewDBConnection)

func MainSingletonExample() {

	results := make(chan *DBConnection)

	for i := 0; i < 10; i++ {
		go func() {
			// First call will create the instance
			instance := DBConnectionSingleton.GetInstance()
			results <- instance
		}()
	}

	for i := 0; i < 10; i++ {
		instance := <-results
		// all the pointers are the same
		fmt.Printf("%p \n", instance)
	}
}
