package ports

import "github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"

type UserServices interface {
	GetUser(user string, pass string) (domain.Users, error)
}

type UserRepository interface {
	GetUser(user string, pass string) (domain.Users, error)
}
