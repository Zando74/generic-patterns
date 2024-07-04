package main

import (
	"fmt"

	"github.com/Zando74/generic-patterns/behavioral"
)

// LOGIN USE CASE

type UserLoginRequestData struct {
	Username string
	Roles    []string
}

type UserLoginResultData struct {
	IsAuth, IsAdmin bool
}

func (u UserLoginResultData) String() string {
	return fmt.Sprintf(" { IsAuth: %v, IsAdmin: %v }", u.IsAuth, u.IsAdmin)
}

type User struct {
	Broker *behavioral.Broker[UserLoginRequestData, UserLoginResultData]
	Name   string
	Roles  []string
}

func NewUser(name string, roles []string, broker *behavioral.Broker[UserLoginRequestData, UserLoginResultData]) *User {
	return &User{Name: name, Roles: roles, Broker: broker}
}

func (u *User) CanAccess() error {
	q := behavioral.Query[UserLoginRequestData, UserLoginResultData]{
		Data:   UserLoginRequestData{Username: u.Name, Roles: u.Roles},
		Result: UserLoginResultData{IsAuth: false, IsAdmin: false},
		Error:  nil,
	}
	u.Broker.Fire(&q)
	return q.Error
}

type isAuthModifier struct {
	behavioral.Handler[UserLoginRequestData, UserLoginResultData]
}

func (authModifier *isAuthModifier) Handle(q *behavioral.Query[UserLoginRequestData, UserLoginResultData]) {
	for _, role := range q.Data.Roles {
		if role == "user" || role == "admin" {
			q.Result.IsAuth = true
			return
		}
	}

	q.Error = fmt.Errorf("UNAUTHENTICATED")
}

type IsAdminModifier struct {
	behavioral.Handler[UserLoginRequestData, UserLoginResultData]
}

func (adminModifier *IsAdminModifier) Handle(q *behavioral.Query[UserLoginRequestData, UserLoginResultData]) {

	for _, role := range q.Data.Roles {
		if role == "admin" {
			q.Result.IsAdmin = true
			return
		}
	}

	q.Error = fmt.Errorf("UNAUTHORIZED")
}

func AuthenticatedRouteCheck(user User) error {
	accessBroker := behavioral.NewBroker[UserLoginRequestData, UserLoginResultData]()
	accessBroker.Subscribe(&isAuthModifier{})
	user.Broker = accessBroker
	return user.CanAccess()

}

func AdminRouteCheck(user User) error {
	accessBroker := behavioral.NewBroker[UserLoginRequestData, UserLoginResultData]()
	accessBroker.Subscribe(&isAuthModifier{})
	accessBroker.Subscribe(&IsAdminModifier{})
	user.Broker = accessBroker
	return user.CanAccess()
}

func Scenario(user User) {
	err := AuthenticatedRouteCheck(user)
	if err != nil {
		fmt.Printf("User %s try to access to Basic route -- Error : %s \n", user.Name, err)
	} else {
		fmt.Printf("User %s try to access to Basic route -- Success \n", user.Name)
	}

	err = AdminRouteCheck(user)
	if err != nil {
		fmt.Printf("User %s try to access to Admin route -- Error : %s \n", user.Name, err)
	} else {
		fmt.Printf("User %s try to access to Admin route -- Success \n", user.Name)
	}

}

func MainChainOfResponsibilityBrokerExample() {

	// User Created from a valid JWT Token
	jack := NewUser("Jack", []string{}, nil)
	john := NewUser("John", []string{"user"}, nil)
	jane := NewUser("Jane", []string{"admin"}, nil)

	// Try to access to differents routes
	Scenario(*jack)
	Scenario(*john)
	Scenario(*jane)

}
