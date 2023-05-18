package services

import (
	"fmt"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
)

type CreateUserUseCase struct {
	userDB database.UserInterface
}

type UserUseCaseInterface interface {
	Create(user *entities.User) error
	FindByEmail(string) (*entities.User, error)
	FindById(string) (*entities.User, error)
	Delete(id string) error
}

func NewUserUseCase(userDB database.UserInterface) UserUseCaseInterface {
	return &CreateUserUseCase{
		userDB: userDB,
	}
}

func (udb *CreateUserUseCase) Create(user *entities.User) error {

	fmt.Println(user)

	err := udb.userDB.Create(user)

	return err

}
func (udb *CreateUserUseCase) FindByEmail(e string) (*entities.User, error) {
	u, err := udb.userDB.FindByEmail(e)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (udb *CreateUserUseCase) FindById(id string) (*entities.User, error) {
	u, err := udb.userDB.FindById(id)

	if err != nil {
		return nil, err
	}

	return u, nil
}
func (udb *CreateUserUseCase) Delete(id string) error {
	err := udb.userDB.Delete(id)

	return err
}
