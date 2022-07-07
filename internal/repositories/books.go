package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"gorm.io/gorm"
)

type books struct {
	DB *gorm.DB
}

func NewBooksRepository(db *gorm.DB) *books {
	return &books{
		db,
	}
}

func (repo *books) GetId(id int) (domain.Books, error) {
	book := domain.Books{}
	err := repo.DB.First(&book, id).Error
	if err != nil {
		return book, errors.New(fmt.Sprintf("Not Found book with ID %d", id))
	}
	return book, nil
}

func (repo *books) Get() ([]domain.Books, error) {
	book := []domain.Books{}
	err := repo.DB.Find(&book).Error
	if err != nil {
		return nil, errors.New("Record is empty")
	}
	return book, nil
}

func (repo *books) BookAdd(book *domain.Books) (domain.Books, error) {
	book.CreatedBy = 1
	book.CreateAt = time.Now()
	book.UpdateAt = time.Now()
	result := repo.DB.Create(&book)
	if result.Error != nil {
		return domain.Books{}, errors.New("Gagal Menambah data")
	}

	resultBook := domain.Books{}
	repo.DB.First(&resultBook, book.ID)

	return resultBook, nil
}

func (repo *books) BookUpdate(book *domain.Books) (domain.Books, error) {
	book.UpdateAt = time.Now()
	result := repo.DB.Model(domain.Books{}).Where("id = ?", book.ID).Updates(book)
	if result.Error != nil {
		return domain.Books{}, result.Error
	}

	resultBook := domain.Books{}
	repo.DB.First(&resultBook, book.ID)
	return resultBook, nil
}

func (repo *books) BookDelete(id int) error {
	result := repo.DB.Where("id=?", id).Delete(domain.Books{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
