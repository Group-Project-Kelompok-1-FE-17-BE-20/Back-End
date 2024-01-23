package data

import (
	"Laptop/app/database"
	"Laptop/features/shoppingcart"
)

func CoreCartToGorm(core shoppingcart.CoreCart) database.ShoppingCart {
	return database.ShoppingCart{
		UserID: core.UserID,
		Status: core.Status,
	}
}

func CartGormToCartCore(model database.ShoppingCart) shoppingcart.CoreCart {
	return shoppingcart.CoreCart{
		UserID: model.UserID,
		Status: model.Status,
	}
}
