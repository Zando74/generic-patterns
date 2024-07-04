package main

import (
	"fmt"

	"github.com/Zando74/generic-patterns/behavioral"
)

type DoctorService struct {
	behavioral.Observer
	Name string
}

func (d DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor of %s has been called for %s \n", d.Name, data.(string))
}

type Person struct {
	behavioral.Observable[DoctorService]
	Name string
}

func NewPerson(name string) *Person {

	return &Person{
		Observable: behavioral.NewObservable[DoctorService](),
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Notify(p.Name)
}

func MainObserverExample() {

	p := NewPerson("John")

	ds1 := &DoctorService{Name: "Hospital 1"}
	ds2 := &DoctorService{Name: "Hospital 2"}
	p.Subscribe(ds1)
	p.Subscribe(ds2)

	p.CatchACold()

}
