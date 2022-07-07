package main

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/core/services"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/handler"
	"github.com/upgradeskill/bookstore-nurfauzanhidayat/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/db_books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	bookRepository := repositories.NewBooksRepository(db)
	bookService := services.NewBookService(bookRepository)
	echoHandler := handler.NewEchoHandler(bookService)

	userRepository := repositories.NewUsersRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	e.GET("/", echoHandler.EchoHelloWorld)
	e.GET("/book", echoHandler.GetBook)
	e.POST("/book", echoHandler.AddBook)
	e.POST("/login", userHandler.GetUser)
	e.PUT("/book", echoHandler.UpdateBook)
	e.DELETE("/book", echoHandler.DeleteBook)

	e.Logger.Fatal(e.Start(":1323"))
}
