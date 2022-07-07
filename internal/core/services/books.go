package services

import (
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/ports"
)

type bookService struct {
	bookRepository ports.BookRepository
}

func NewBookService(bookRepository ports.BookRepository) *bookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (srv *bookService) GetId(id int) (domain.Books, error) {
	book, err := srv.bookRepository.GetId(id)
	if err != nil {
		return domain.Books{}, err
	}
	return book, nil
}

func (srv *bookService) Get() ([]domain.Books, error) {
	book, err := srv.bookRepository.Get()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (srv *bookService) BookAdd(bookRequest *domain.Books) (domain.Books, error) {
	book, err := srv.bookRepository.BookAdd(bookRequest)
	if err != nil {
		return domain.Books{}, err
	}
	return book, nil
}

func (srv *bookService) BookUpdate(bookRequest *domain.Books) (domain.Books, error) {
	book, err := srv.bookRepository.BookUpdate(bookRequest)
	if err != nil {
		return domain.Books{}, err
	}
	return book, nil
}

func (srv *bookService) BookDelete(id int) error {
	err := srv.bookRepository.BookDelete(id)
	if err != nil {
		return err
	}
	return nil
}
