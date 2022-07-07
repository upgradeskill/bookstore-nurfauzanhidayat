package ports

import "github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"

type BookServices interface {
	GetId(id int) (domain.Books, error)
	BookAdd(book *domain.Books) (domain.Books, error)
	BookUpdate(book *domain.Books) (domain.Books, error)
	BookDelete(id int) error
	Get() ([]domain.Books, error)
}

type BookRepository interface {
	GetId(id int) (domain.Books, error)
	BookAdd(book *domain.Books) (domain.Books, error)
	BookUpdate(book *domain.Books) (domain.Books, error)
	BookDelete(id int) error
	Get() ([]domain.Books, error)
}
