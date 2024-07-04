package main

import "fmt"

type Customer struct {
	Name          string
	Level         int
	ProblemSolved bool
}

type BaseSupportHandler interface {
	Handle(*Customer)
	SetNext(BaseSupportHandler)
}

type BaseSupport struct {
	Next BaseSupportHandler
}

func (ch *BaseSupport) SetNext(next BaseSupportHandler) {
	ch.Next = next
}

func (bs *BaseSupport) Handle(customer *Customer) {
	if customer.Level == 1 {
		fmt.Println("Problem solved by lvl 1 for customer", customer.Name)
		customer.ProblemSolved = true
	} else if bs.Next != nil {
		bs.Next.Handle(customer)
	}
}

type TechnicalSupport struct {
	Next BaseSupportHandler
}

func (ch *TechnicalSupport) SetNext(next BaseSupportHandler) {
	ch.Next = next
}

func (ts *TechnicalSupport) Handle(customer *Customer) {
	if customer.Level == 2 {
		fmt.Println("Problem solved by lvl 2 for customer", customer.Name)
		customer.ProblemSolved = true
	} else if ts.Next != nil {
		ts.Next.Handle(customer)
	}
}

type AdvancedSupport struct {
	Next BaseSupportHandler
}

func (ch *AdvancedSupport) SetNext(next BaseSupportHandler) {
	ch.Next = next
}

func (as *AdvancedSupport) Handle(customer *Customer) {
	if customer.Level == 3 {
		fmt.Println("Problem solved by lvl 3 for customer", customer.Name)
		customer.ProblemSolved = true
	} else if as.Next != nil {
		as.Next.Handle(customer)
	}
}

func MainChainOfResponsibility() {
	customer := &Customer{
		Name:  "John",
		Level: 3,
	}

	baseSupport := &BaseSupport{}
	technicalSupport := &TechnicalSupport{}
	advancedSupport := &AdvancedSupport{}

	baseSupport.SetNext(technicalSupport)
	technicalSupport.SetNext(advancedSupport)

	baseSupport.Handle(customer)
}
