package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/domain"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/ports"
)

type EchoHandler struct {
	bookService ports.BookServices
}

func NewEchoHandler(bookService ports.BookServices) *EchoHandler {
	return &EchoHandler{
		bookService: bookService,
	}
}

func (hand *EchoHandler) EchoHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello From Echo")
}

func (hand *EchoHandler) GetBook(c echo.Context) error {
	resp := make(map[string]interface{})
	if c.QueryParam("id") != "" {
		id, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			resp["message"] = "Not Valid ID"
			return c.JSON(http.StatusBadRequest, resp)
		}

		book, err := hand.bookService.GetId(id)
		if err != nil {
			resp["message"] = err.Error()
			return c.JSON(http.StatusNotFound, resp)
		}

		resp["message"] = "Data Found"
		resp["data"] = book
		return c.JSON(http.StatusOK, resp)

	} else {
		book, err := hand.bookService.Get()
		if err != nil {
			resp["message"] = err.Error()
			return c.JSON(http.StatusNotFound, resp)
		}

		resp["message"] = "Data Found"
		resp["data"] = book
		return c.JSON(http.StatusOK, resp)
	}
}

func (hand *EchoHandler) AddBook(c echo.Context) error {
	resp := make(map[string]interface{})
	book := domain.Books{}

	err := c.Bind(&book)
	if err != nil {
		log.Println(err)
		resp["message"] = "Failed to parsing data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	result, err := hand.bookService.BookAdd(&book)

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp["message"] = "Success"
	resp["data"] = result
	return c.JSON(http.StatusCreated, resp)
}

func (hand *EchoHandler) UpdateBook(c echo.Context) error {
	resp := make(map[string]interface{})
	book := domain.Books{}

	err := c.Bind(&book)
	if err != nil {
		log.Println(err)
		resp["message"] = "Failed to parsing data"
		return c.JSON(http.StatusBadRequest, resp)
	}

	result, err := hand.bookService.BookUpdate(&book)

	if err != nil {
		resp["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp["message"] = "Success"
	resp["data"] = result
	return c.JSON(http.StatusCreated, resp)
}

func (hand *EchoHandler) DeleteBook(c echo.Context) error {
	resp := make(map[string]interface{})
	if c.QueryParam("id") != "" {
		id, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			resp["message"] = "Not Valid ID"
			return c.JSON(http.StatusBadRequest, resp)
		}

		errResult := hand.bookService.BookDelete(id)
		if errResult != nil {
			resp["message"] = err.Error()
			return c.JSON(http.StatusNotFound, resp)
		}

		resp["message"] = "Data Delete"
		return c.JSON(http.StatusOK, resp)

	} else {
		resp["message"] = "Id is requried"
		return c.JSON(http.StatusBadRequest, resp)
	}
}
