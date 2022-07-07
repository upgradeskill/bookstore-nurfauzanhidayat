package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (srv *userService) GetUser(username string, password string) (domain.Users, error) {
	user, err := srv.userRepository.GetUser(username, password)

	userClaims := domain.Users{}
	userClaims.Userid = user.Userid
	userClaims.Profile = user.Profile

	claims := &domain.UserClaims{
		user.Userid,
		user.Profile,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return domain.Users{}, err
	}
	user.Token = t
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}
