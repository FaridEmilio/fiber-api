package user

import (
	"errors"

	"github.com/faridEmilio/fiber-api/pkg/dtos/userdtos"
)

// Defino la interfaz: Digo todo lo que el servicio puede hacer
type UserService interface {
	PostCreateUserService(request userdtos.RequestPostUser) (status bool, erro error)
}

// Defino estructura
type userService struct {
	repository UserRepository
	//dbHelpersRepository dbhelpers.DbHelpersRepository
	//store               storage.Storage // un objeto que implementa un metodo de almacenamiento
	//commonsService      commons.Commons // otras funciones de ayuda
}

// Defino un constructor
func NewUserService(r UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

// CreateUserService implements UserService.
func (s *userService) PostCreateUserService(request userdtos.RequestPostUser) (status bool, erro error) {
	// solicitar repository
	erro = s.repository.PostCreateUserRepository(request.ToEntity(false))

	if erro != nil {
		status = false
		erro = errors.New(erro.Error())
	}

	status = true
	return
}
