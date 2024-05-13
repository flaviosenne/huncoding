package view

import (
	"github.com/flaviosenne/huncoding/src/controller/model/response"
	"github.com/flaviosenne/huncoding/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
