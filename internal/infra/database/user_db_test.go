package database_test

import (
	"testing"

	"github.com/IcaroSilvaFK/goexpert_first_api/internal/entities"
	"github.com/IcaroSilvaFK/goexpert_first_api/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldCreateNewUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.User{})
	userDB := database.NewUserDB(db)

	user, _ := entities.NewUserEntity("jog@j.com", "Jhon", "123456")
	err = userDB.Create(user)

	assert.Nil(t, err)

	var u entities.User

	tx := db.Find(&u, "id = ?", user.ID)

	assert.Nil(t, tx.Error)
	assert.Equal(t, u.ID, user.ID)
	assert.NotNil(t, user.Password)
}

func TestShouldFinUserByEmail(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)

	db.AutoMigrate(&entities.User{})
	userDB := database.NewUserDB(db)

	user, _ := entities.NewUserEntity("jog@j.com", "Jhon", "123456")
	err = userDB.Create(user)

	assert.Nil(t, err)

	u, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Email, user.Email)
	assert.Equal(t, u.ID, user.ID)
	assert.NotNil(t, u.Password)

}

func TestShouldFinUserById(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)

	db.AutoMigrate(&entities.User{})
	userDB := database.NewUserDB(db)

	user, _ := entities.NewUserEntity("jog@j.com", "Jhon", "123456")
	err = userDB.Create(user)

	assert.Nil(t, err)

	u, err := userDB.FindById(user.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Email, user.Email)
	assert.Equal(t, u.ID, user.ID)
	assert.NotNil(t, u.Password)

}

func TestShouldDeleteUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	assert.Nil(t, err)

	db.AutoMigrate(&entities.User{})
	userDB := database.NewUserDB(db)

	user, _ := entities.NewUserEntity("jog@j.com", "Jhon", "123456")
	err = userDB.Create(user)

	assert.Nil(t, err)

	err = userDB.Delete(user.ID.String())
	assert.Nil(t, err)

	u, err := userDB.FindById(user.ID.String())
	assert.Error(t, err)

	assert.Nil(t, u)
}
