package user

import (
	"errors"

	"github.com/faridEmilio/fiber-api/internal/database"
	"github.com/faridEmilio/fiber-api/pkg/entities"
	"gorm.io/gorm/clause"
)

// Defino una interface
type UserRepository interface {
	PostCreateUserRepository(entityUser entities.User) (err error)
	GetUsersRepository()
	GetUserByIdRepository(userId uint) (entityUser entities.User, rowAffected bool, err error)
	UpdateUserRepository(entityUser entities.User) (err error)
	DeleteUserRepository()
}

func NewUserRepository(db *database.DbInstance) UserRepository {
	return &userRepository{
		SqlClient: db,
	}
}

type userRepository struct {
	SqlClient *database.DbInstance
}

func (r *userRepository) PostCreateUserRepository(entityUser entities.User) (err error) {

	resp := r.SqlClient.Omit(clause.Associations).Create(&entityUser)

	if resp.Error != nil {
		err = errors.New("Error al crear un usuario")
	}
	return
}

// DeleteUserRepository implements UserRepository.
func (*userRepository) DeleteUserRepository() {
	panic("unimplemented")
}

// GetUserByIdRepository implements UserRepository.
func (*userRepository) GetUserByIdRepository(userId uint) (entityUser entities.User, rowAffected bool, err error) {
	panic("unimplemented")
}

// GetUsersRepository implements UserRepository.
func (*userRepository) GetUsersRepository() {
	panic("unimplemented")
}

// UpdateUserRepository implements UserRepository.
func (*userRepository) UpdateUserRepository(entityUser entities.User) (err error) {
	panic("unimplemented")
}

func findUser(id int, user *pkg.User) error {
	database.DbInstance.Find(&user, "id= ?", id)

	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}
