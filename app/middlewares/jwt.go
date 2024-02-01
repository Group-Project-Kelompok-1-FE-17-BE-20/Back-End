package middlewares

import (
	config "Laptop/app/configs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var appConfig = config.ReadEnv()

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(appConfig.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId uint /*, userRole string*/) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	// claims["userRole"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(appConfig.JWT_SECRET))

}

func ExtractTokenUserId(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}

//	func ExtractTokenUserRole(e echo.Context) string {
//		user := e.Get("user").(*jwt.Token)
//		if user.Valid {
//			claims := user.Claims.(jwt.MapClaims)
//			userRole := claims["userRole"].(string)
//			return userRole
//		}
//		return ""
//	}
func ExtractToken(e echo.Context) (float64, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := claims["userId"].(float64)
		//username := claims["username"].(string)
		return id, nil //username
	}
	return 0, errors.New("failed to extract jwt-token")
}
