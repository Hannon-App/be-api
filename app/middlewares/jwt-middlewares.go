package middlewares

import (
	"Hannon-app/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT_SECRRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId uint, adminId uint, tenantId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["admin"] = adminId
	claims["tenant"] = tenantId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
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

func ExtractTokenAdminId(e echo.Context) uint {
	admin := e.Get("admin").(*jwt.Token)
	if admin.Valid {
		claims := admin.Claims.(jwt.MapClaims)
		adminId := claims["adminId"].(float64)
		return uint(adminId)
	}
	return 0
}

func ExtractTokenTenantId(e echo.Context) uint {
	tenant := e.Get("tenant").(*jwt.Token)
	if tenant.Valid {
		claims := tenant.Claims.(jwt.MapClaims)
		tenantId := claims["tenantId"].(float64)
		return uint(tenantId)
	}
	return 0
}
