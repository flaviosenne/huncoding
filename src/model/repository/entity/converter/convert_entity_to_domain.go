package converter

import (
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository/entity"
)

func ConverterEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetID(entity.ID.Hex())
	return domain
}
