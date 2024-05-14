package converter

import (
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository/entity"
)

func ConverterDomainToEntity(model model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    model.GetEmail(),
		Name:     model.GetName(),
		Password: model.GetPassword(),
		Age:      model.GetAge(),
	}
}
