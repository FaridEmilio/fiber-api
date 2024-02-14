package user

import (
	"github.com/faridEmilio/fiber-api/pkg/entities"
)

// Defino una interface
type UserRepository interface {
	CreateUserRepository(entityUser entities.User) (err error)
	GetUsersRepository()
	GetUserByIdRepository(userId uint) (entityUser entities.User, rowAffected bool, err error)
	UpdateUserRepository(entityUser entities.User) (err error)
	DeleteUserRepository()
}
