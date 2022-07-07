package handler

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/ports"
)

type UserHandler struct {
	userService ports.UserServices
}

func NewUserHandler(userService ports.UserServices) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (hand *UserHandler) GetUser(c echo.Context) error {
	resp := make(map[string]interface{})
	user := domain.Users{}

	err := c.Bind(&user)
	if err != nil {
		resp["message"] = "Failed to parsing data"
		return c.JSON(http.StatusBadRequest, resp)
	}
	password := md5.Sum([]byte(user.Password))
	result, err := hand.userService.GetUser(user.Username, hex.EncodeToString(password[:]))

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	userResponse := make(map[string]string)
	userResponse["username"] = result.Username
	userResponse["profile"] = result.Profile
	userResponse["token"] = result.Token

	resp["message"] = "Success"
	resp["data"] = userResponse
	return c.JSON(http.StatusCreated, resp)
}
