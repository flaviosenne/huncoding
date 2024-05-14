package model

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int8

	SetID(string)

	GetJsonValue() (string, error)

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
