package router

import (
	"Laptop/app/middlewares"

	_userRepo "Laptop/features/user/data"
	_userHandler "Laptop/features/user/handler"
	_userService "Laptop/features/user/service"

	_dataProduct "Laptop/features/product/data"
	_productHandler "Laptop/features/product/handler"
	_productService "Laptop/features/product/service"

	_StoreRepo "Laptop/features/store/data"
	_StoreHandler "Laptop/features/store/handler"
	_StoreService "Laptop/features/store/service"

	_CartData "Laptop/features/shoppingcart/data"
	_CartHandler "Laptop/features/shoppingcart/handler"
	_CartService "Laptop/features/shoppingcart/service"

	_ItemData "Laptop/features/shoppingcartitem/data"
	_ItemHandler "Laptop/features/shoppingcartitem/handler"
	_ItemService "Laptop/features/shoppingcartitem/service"

	_adminRepo "Laptop/features/admin/data"
	_adminHandler "Laptop/features/admin/handler"
	_adminService "Laptop/features/admin/service"

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

	StoreRepo := _StoreRepo.New(db)
	StoreService := _StoreService.New(StoreRepo)
	StoreHandler := _StoreHandler.New(StoreService)

	cartData := _CartData.New(db)
	cartService := _CartService.New(cartData)
	cartHandlerAPI := _CartHandler.New(cartService)

	itemData := _ItemData.New(db)
	itemService := _ItemService.New(itemData)
	itemHandlerAPI := _ItemHandler.New(itemService)

	adminRepo := _adminRepo.New(db)
	adminService := _adminService.New(adminRepo)
	adminHandlerAPI := _adminHandler.New(adminService)

	// user
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUser)
	e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUserById, middlewares.JWTMiddleware())

	// store
	// e.GET("/stores", StoreHandler.GetAllStore,middlewares.JWTMiddleware())
	e.GET("/stores/:store_id", StoreHandler.GetStoreById, middlewares.JWTMiddleware())
	e.POST("/stores", StoreHandler.CreateStore, middlewares.JWTMiddleware())
	e.PUT("/stores/:store_id", StoreHandler.UpdateStoreById, middlewares.JWTMiddleware())
	e.DELETE("/stores/:store_id", StoreHandler.DeleteStoreById, middlewares.JWTMiddleware())

	// product
	e.POST("/products", productHandlerAPI.CreateProduct, middlewares.JWTMiddleware())
	e.PUT("/products/:product_id", productHandlerAPI.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:product_id", productHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProducts, middlewares.JWTMiddleware())
	e.GET("/products/:product_id", productHandlerAPI.GetSingleProduct)
	// e.GET("products/:username", productHandlerAPI.GetProductofUser)

	// shopping cart
	e.GET("/shopping-cart", cartHandlerAPI.GetCart, middlewares.JWTMiddleware())

	// shopping cart item
	e.POST("/shopping-cart", itemHandlerAPI.CreateItem, middlewares.JWTMiddleware())
	e.PUT("/shopping-cart", itemHandlerAPI.UpdateItem, middlewares.JWTMiddleware())
	//e.DELETE("/shopping-cart", itemHandlerAPI.Delete, middlewares.JWTMiddleware())

	// admin
	e.POST("/logins", adminHandlerAPI.Login)
	// e.GET("/users", adminHandlerAPI.GetAllUser)
	e.GET("/userss", userHandlerAPI.GetAllUser)
}
