package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	id       string
	name     string
	email    string
	password string
	age      int8
}

func (ud *userDomain) GetJsonValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}
