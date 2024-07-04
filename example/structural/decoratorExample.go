package main

import (
	"fmt"
	"generic-patterns/structural"
)

type Shape interface {
	Render() string
	Display() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f, ",
		c.Radius)
}

func (c *Circle) Display() string {
	return fmt.Sprintf("Circle of radius %f, ",
		c.Radius)
}

type ColoredShape struct {
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("has the color %s, ", c.Color)
}

func (c *ColoredShape) Display() string {
	return fmt.Sprintf("has the color %s, ", c.Color)
}

type TransparentShape struct {
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("has %f%% transparency.", t.Transparency*100.0)
}

func (t *TransparentShape) Display() string {
	return fmt.Sprintf("has %f%% transparency.", t.Transparency*100.0)
}

func MainDecoractorExample() {

	// Decorator execution behavior, will be executed for each wrapped items
	renderHandler := func(s Shape) {
		// You should use any Decorable interface method here
		fmt.Println(s.Render())
	}

	structural.
		// Initialize a new Decorator for Shape and Wrap the Circle struct
		NewDecorator[Shape](&Circle{Radius: 2}).
		// Wrap the ColoredShape struct over the Circle
		Wrap(&ColoredShape{Color: "red"}).
		// Wrap the TransparentShape struct over the ColoredShape
		Wrap(&TransparentShape{Transparency: 0.5}).
		// Setting Shape.Render() implementation to be executed for each wrapped items
		SetExecutionHandler(&renderHandler).
		// Output: Circle of radius 2, has the color red, has 50% transparency.
		Execute()

	err := structural.
		// Initialize a new Decorator for Shape and Wrap the Circle struct
		NewDecorator[Shape](&Circle{Radius: 2}).
		// Trying to execute without setting an execution handler
		Execute()

	if err != nil {
		// Output: You should must set an execution handler before Executing it
		fmt.Println(err)
	}

	// You can use Factory Generator to Generate Custom Decorators Factories and avoid error risks

	displayHandler := func(s Shape) {
		fmt.Println(s.Display())
	}

	coloredTransparentCircleDisplayFactory := structural.GenerateDecoratorFactory[Shape](
		// Decorator execution behavior
		&displayHandler,
		// Initial Decorable struct
		&Circle{Radius: 2},
		&ColoredShape{Color: "red"},
		&TransparentShape{Transparency: 0.5},
		// ... Decorables to be wrapped in order (You can specify order you need in case of state sharing)
	)

	coloredTransparentCircleDisplay := coloredTransparentCircleDisplayFactory.Make()
	// Output: Circle of radius 2, has the color red, has 50% transparency.
	coloredTransparentCircleDisplay.Execute()

}
