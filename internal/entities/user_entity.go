package entities

import (
	"errors"
	"regexp"

	"github.com/IcaroSilvaFK/goexpert_first_api/pkg/entity"
	"github.com/IcaroSilvaFK/goexpert_first_api/pkg/utils"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUserEntity(email, name, password string) (*User, error) {

	p, err := utils.GenerateHash(password)

	if err != nil {
		return nil, err
	}

	if err = verifyEmail(email); err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.GenerateID(),
		Email:    email,
		Name:     name,
		Password: p,
	}, nil

}

func (u *User) Save() error {
	return nil
}
func (u *User) Update(name, email string) error {
	return nil
}
func (u *User) Delete(id string) error {
	return nil
}

func (u *User) ValidatePassword(p string) error {

	err := utils.VerifyHash(p, u.Password)

	return err
}

func verifyEmail(email string) error {

	pattern := `[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`

	regex := regexp.MustCompile(pattern)

	if !regex.Match([]byte(email)) {
		return errors.New("invalid email")
	}

	return nil
}
