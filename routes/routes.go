package routes

import (
	"erp/constants"
	"erp/controllers"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	eUser := e.Group("users")
	eUser.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eUser.GET("", controllers.GetUsersController)
	eUser.GET("/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	eUser.DELETE("/:id", controllers.DeleteUserController)
	eUser.PUT("/:id", controllers.UpdateUserController)

	eOfficer := e.Group("officers")
	eOfficer.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eOfficer.GET("", controllers.GetOfficersController)
	eOfficer.GET("/:id", controllers.GetOfficerController)
	eOfficer.POST("", controllers.CreateOfficerController)
	eOfficer.DELETE("/:id", controllers.DeleteOfficerController)
	eOfficer.PUT("/:id", controllers.UpdateOfficerController)

	eGudang := e.Group("gudangs")
	eGudang.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eGudang.GET("", controllers.GetGudangsController)
	eGudang.GET("/:id", controllers.GetGudangController)
	eGudang.POST("", controllers.CreateGudangController)
	eGudang.DELETE("/:id", controllers.DeleteGudangController)
	eGudang.PUT("/:id", controllers.UpdateGudangController)

	ePersediaan := e.Group("persediaan_barang")
	ePersediaan.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	ePersediaan.GET("", controllers.GetPersediaansController)
	ePersediaan.GET("/:id", controllers.GetPersediaanController)
	ePersediaan.POST("", controllers.CreatePersediaanController)
	ePersediaan.DELETE("/:id", controllers.DeletePersediaanController)
	ePersediaan.PUT("/:id", controllers.UpdatePersediaanController)

	eBarang := e.Group("barang_masuk")
	eBarang.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBarang.GET("", controllers.GetBarangsController)
	eBarang.GET("/:id", controllers.GetBarangController)
	eBarang.POST("", controllers.CreateBarangController)
	eBarang.DELETE("/:id", controllers.DeleteBarangController)
	eBarang.PUT("/:id", controllers.UpdateBarangController)

	eBom := e.Group("bom")
	eBom.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eBom.GET("", controllers.GetBOMsController)
	eBom.GET("/:id", controllers.GetBOMController)
	eBom.POST("", controllers.CreateBOMController)
	eBom.DELETE("/:id", controllers.DeleteBOMController)
	eBom.PUT("/:id", controllers.UpdateBOMController)

	eWo := e.Group("wo")
	eWo.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eWo.GET("", controllers.GetWOsController)
	eWo.GET("/:id", controllers.GetWOController)
	eWo.POST("", controllers.CreateWOController)
	eWo.DELETE("/:id", controllers.DeleteWOController)
	eWo.PUT("/:id", controllers.UpdateWOController)

	eProduk := e.Group("produk")
	eProduk.Use(mid.JWT([]byte(constants.SECRET_KEY)))

	eProduk.GET("", controllers.GetProduksController)
	eProduk.GET("/:id", controllers.GetProdukController)
	eProduk.POST("", controllers.CreateProdukController)
	eProduk.DELETE("/:id", controllers.DeleteProdukController)
	eProduk.PUT("/:id", controllers.UpdateProdukController)

	e.POST("/login", controllers.LoginController)

	return e
}
