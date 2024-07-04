package main

type Account struct{}

func (a *Account) Create() error {
	println("Account Created")
	return nil
}

type User struct{}

func (u *User) Create() error {
	println("User Created")
	return nil
}

type Role struct{}

func (r *Role) Create() error {
	println("Role Created")
	return nil
}

type Permission struct{}

func (p *Permission) Create() error {
	println("Permission Created")
	return nil
}

type Registration struct {
	Account    *Account
	User       *User
	Role       *Role
	Permission *Permission
}

func NewRegistration() *Registration {
	return &Registration{
		Account:    &Account{},
		User:       &User{},
		Role:       &Role{},
		Permission: &Permission{},
	}
}

func (r *Registration) InscriptionUseCase() error {
	err := r.Account.Create()

	if err != nil {
		return err
	}

	err = r.User.Create()

	if err != nil {
		return err
	}

	err = r.Role.Create()

	if err != nil {
		return err
	}

	err = r.Permission.Create()

	if err != nil {
		return err
	}

	return nil
}

func MainFacadeExample() {
	registrationFacade := NewRegistration()

	err := registrationFacade.InscriptionUseCase()

	if err != nil {
		println(err.Error())
	}

}
