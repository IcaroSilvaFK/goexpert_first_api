package database

import (
	"fmt"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *User {

	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entities.User) error {
	fmt.Println(user)
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(e string) (*entities.User, error) {

	var user entities.User

	tx := u.DB.Where("email = ?", e).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (u *User) FindById(id string) (*entities.User, error) {
	var user entities.User

	tx := u.DB.Where("id = ?", id).First(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil

}

func (u *User) Delete(id string) error {

	tx := u.DB.Delete(&entities.User{}, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
