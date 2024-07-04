package main

import (
	"fmt"
	"generic-patterns/creational"
)

type Car struct {
	Brand  string
	Model  string
	Option string
}

type CarBuilder struct {
	// A builder using a functionnal approach
	// Object is created only when Build() is called
	// store setters as functions to call successively later
	// Building rules should be updated dynamically
	creational.FunctionalBuilder[Car]
}

func (builder *CarBuilder) SetBrand(brand string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Brand = brand
	})
	return builder
}

func (builder *CarBuilder) SetModel(model string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Model = model
	})
	return builder
}

func (builder *CarBuilder) SetOption(option string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Option = option
	})
	return builder
}

func GenerateCarBuilder(carType string) *CarBuilder {

	if carType == "AudiR8" {
		return (&CarBuilder{}).
			SetBrand("Audi").
			SetModel("R8").
			SetOption("V10")
	}
	if carType == "CitroenC3" {
		return (&CarBuilder{}).
			SetBrand("Citroen").
			SetModel("C3")
	}
	return nil
}

func MainFunctionnalBuilder() {

	audiR8CarBuilder := GenerateCarBuilder("AudiR8")
	citroenC3CarBuilder := GenerateCarBuilder("CitroenC3")

	c1 := audiR8CarBuilder.Build()
	c2 := citroenC3CarBuilder.Build()

	// air conditioning available, update the builder
	citroenC3CarBuilder.SetOption("air conditioning")

	// Next cars will have air conditioning
	c3 := citroenC3CarBuilder.Build()

	// Reset the builder
	citroenC3CarBuilder.Reset()

	// Next car will have no brand, model and option
	c4 := citroenC3CarBuilder.Build()

	fmt.Println(c1) // Output: &{Audi R8 V10}
	fmt.Println(c2) // Output: &{Citroen C3 ""}
	fmt.Println(c3) // Output: &{Citroen C3 air conditioning}
	fmt.Println(c4) // Output: &{"" "" ""}

}
