package main

import (
	"fmt"
	"generic-patterns/structural"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func NewRed() *Color {
	return &Color{R: 255, G: 0, B: 0}
}

func NewGreen() *Color {
	return &Color{R: 0, G: 255, B: 0}
}

func NewBlue() *Color {
	return &Color{R: 0, G: 0, B: 255}
}

type Color struct {
	structural.IntrinsicState
	R, G, B uint8
}

type OptimizedShape struct {
	Color *Color
	// ...
}

func NewColorFlyweight() *structural.Flyweight[Color] {

	colorFlyweight := structural.NewFlyweight[Color]()

	colorFlyweight.NewCreationHandler(RED, NewRed)
	colorFlyweight.NewCreationHandler(GREEN, NewGreen)
	colorFlyweight.NewCreationHandler(BLUE, NewBlue)

	return colorFlyweight
}

func MainFlyghtWeighExample() {
	OptimizedCircleFlyweight := NewColorFlyweight()

	OptimizedRedShape := &OptimizedShape{
		Color: OptimizedCircleFlyweight.GetInstance(RED),
	}

	SecondOptimizedRedShape := &OptimizedShape{
		Color: OptimizedCircleFlyweight.GetInstance(RED),
	}

	fmt.Printf("first color address: %p, second color address: %p , Color : %v",
		OptimizedRedShape.Color,
		SecondOptimizedRedShape.Color,
		*OptimizedRedShape.Color,
	) // Should print the same address for both shapes

}
