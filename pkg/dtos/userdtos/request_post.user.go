package userdtos

import "github.com/faridEmilio/fiber-api/pkg/entities"

type RequestPostUser struct {
	ID uint
	//CreatedAt time.Time
	FirstName string
	LastName  string
}

func (rpu *RequestPostUser) ToEntity(isUpdate bool) (e entities.User) {
	// si es un  update se necesita el id de la entidad
	if isUpdate {
		e.ID = rpu.ID
	}

	e.FirstName = rpu.FirstName
	e.LastName = rpu.LastName

	return
}
