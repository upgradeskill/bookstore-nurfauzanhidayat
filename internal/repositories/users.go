package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"gorm.io/gorm"
)

type users struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *users {
	return &users{
		db,
	}
}

func (repo *users) GetUser(username string, password string) (domain.Users, error) {
	user := domain.Users{}
	log.Println(user)
	err := repo.DB.Where("username=? and password=?", username, password).First(&user).Error
	if err != nil {
		return domain.Users{}, errors.New(fmt.Sprintf("Username Not Found"))
	}
	return user, nil
}
