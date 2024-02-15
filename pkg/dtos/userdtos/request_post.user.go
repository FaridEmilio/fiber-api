package userdtos

import "github.com/faridEmilio/fiber-api/pkg/entities"

type RequestPostUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (rpu *RequestPostUser) ToEntity(isUpdate bool) (e entities.User) {
	e.FirstName = rpu.FirstName
	e.LastName = rpu.LastName
	return
}
