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
