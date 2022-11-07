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

	e.GET("/gudangs", controllers.GetGudangsController)
	e.GET("/gudangs/:id", controllers.GetGudangController)
	e.POST("/gudangs", controllers.CreateGudangController)
	e.DELETE("/gudangs/:id", controllers.DeleteGudangController)
	e.PUT("/gudangs/:id", controllers.UpdateGudangController)

	e.GET("/persediaan_barang", controllers.GetPersediaansController)
	e.GET("/persediaan_barang/:id", controllers.GetPersediaanController)
	e.POST("/persediaan_barang", controllers.CreatePersediaanController)
	e.DELETE("/persediaan_barang/:id", controllers.DeletePersediaanController)
	e.PUT("/persediaan_barang/:id", controllers.UpdatePersediaanController)

	e.GET("/barang_masuk", controllers.GetBarangsController)
	e.GET("/barang_masuk/:id", controllers.GetBarangController)
	e.POST("/barang_masuk", controllers.CreateBarangController)
	e.DELETE("/barang_masuk/:id", controllers.DeleteBarangController)
	e.PUT("/barang_masuk/:id", controllers.UpdateBarangController)

	e.GET("/bom", controllers.GetBOMsController)
	e.GET("/bom/:id", controllers.GetBOMController)
	e.POST("/bom", controllers.CreateBOMController)
	e.DELETE("/bom/:id", controllers.DeleteBOMController)
	e.PUT("/bom/:id", controllers.UpdateBOMController)

	e.GET("/wo", controllers.GetWOsController)
	e.GET("/wo/:id", controllers.GetWOController)
	e.POST("/wo", controllers.CreateWOController)
	e.DELETE("/wo/:id", controllers.DeleteWOController)
	e.PUT("/wo/:id", controllers.UpdateWOController)

	e.GET("/produk", controllers.GetProduksController)
	e.GET("/produk/:id", controllers.GetProdukController)
	e.POST("/produk", controllers.CreateProdukController)
	e.DELETE("/produk/:id", controllers.DeleteProdukController)
	e.PUT("/produk/:id", controllers.UpdateProdukController)

	e.POST("/login", controllers.LoginController)

	// eJwtAuth := e.Group("auth")
	// eJwtAuth.Use(middleware.checkAdmin())
	// eJwtAuth.GET("/wo", controllers.GetWOsController)
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
