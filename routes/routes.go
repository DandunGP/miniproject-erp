package routes

import (
	"erp/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/officers", controllers.GetOfficersController)
	e.GET("/officers/:id", controllers.GetOfficerController)
	e.POST("/officers", controllers.CreateOfficerController)
	e.DELETE("/officers/:id", controllers.DeleteOfficerController)
	e.PUT("/officers/:id", controllers.UpdateOfficerController)

	// e.POST("/login", controllers.LoginController)

	// eJwtAuth := e.Group("auth")
	// eJwtAuth.GET("/users/:id", controllers.GetUserController)
	// eJwtAuth.GET("/books", controllers.GetBooksController)
	// eJwtAuth.GET("/books/:id", controllers.GetBookController)

	// eJwtAuth.Use(mid.JWT([]byte(constants.SECRET_KEY)))
	// eJwtAuth.GET("/users", controllers.GetUsersController)
	// eJwtAuth.POST("/users", controllers.CreateUserController)
	// eJwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	// eJwtAuth.PUT("/users/:id", controllers.UpdateUserController)

	// eJwtAuth.POST("/books", controllers.CreateBookController)
	// eJwtAuth.DELETE("/books/:id", controllers.DeleteBookController)
	// eJwtAuth.PUT("/books/:id", controllers.UpdateBookController)
	return e
}

/*
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
*/
