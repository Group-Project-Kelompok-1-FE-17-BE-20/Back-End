package router

import (
	"Laptop/app/middlewares"
	//_projectRepo "Laptop/controllers/project/data"
	//_projectHandler "Laptop/controllers/project/handler"
	//_projectService "Laptop/controllers/project/service"
	//_taskRepo "Laptop/controllers/task/data"
	//_taskHandler "Laptop/controllers/task/handler"
	//_taskService "Laptop/controllers/task/service"
	_userRepo "Laptop/features/user/data"
	_userHandler "Laptop/features/user/handler"
	_userService "Laptop/features/user/service"

	_dataProduct "Laptop/features/product/data"
	_productHandler "Laptop/features/product/handler"
	_productService "Laptop/features/product/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := _userRepo.New(db)
	userService := _userService.New(userRepo)
	userHandlerAPI := _userHandler.New(userService)

	productData := _dataProduct.New(db)
	productService := _productService.New(productData)
	productHandlerAPI := _productHandler.New(productService)

	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUser)
	e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// toko
	// e.GET("/users/:user_id/store-attributes", userHandlerAPI.GetTokoById) // get data toko
	// e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	// e.PUT("/users/:user_id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	// e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// product
	//e.GET("/products", productHandlerAPI.GetAllProducts)
	e.POST("/products", productHandlerAPI.CreateProduct)
	// e.GET("/products/:product_id", productHandlerAPI.GetSingleProduct)
	e.PUT("/products/:product_id", productHandlerAPI.UpdateProduct)
	// e.DELETE("/products/:product_id", productHandlerAPI.Delete)
	// e.GET("products/:username", productHandlerAPI.GetProductofUser)
}
