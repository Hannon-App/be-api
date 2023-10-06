package middlewares

import (
	"Hannon-app/app/config"
	"errors"
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

func CreateTokenUser(userid uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userid
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
}

func CreateTokenTenant(tenantid uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = tenantid
	claims["role"] = "tenant"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
}

func CreateTokenAdmin(adminid uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = adminid
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_SECRRET))
}

func ExtractTokenUser(e echo.Context) (uint, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
		role := claims["role"].(string)
		if role != "user" {
			return 0, errors.New("only user can access")
		}
		return uint(userId), nil
	}
	return 0, errors.New("token invalid")
}

func ExtractTokenTenant(e echo.Context) (uint, error) {
	tenant := e.Get("user").(*jwt.Token)
	if tenant.Valid {
		claims := tenant.Claims.(jwt.MapClaims)
		tenantid := claims["id"].(float64)
		role := claims["role"].(string)
		if role != "tenant" {
			return 0, errors.New("only tenant can access")
		}
		return uint(tenantid), nil
	}
	return 0, errors.New("token invalid")
}

func ExtractTokenAdmin(e echo.Context) (uint, error) {
	admin := e.Get("user").(*jwt.Token)
	if admin.Valid {
		claims := admin.Claims.(jwt.MapClaims)
		adminid := claims["id"].(float64)
		role := claims["role"].(string)
		if role != "admin" {
			return 0, errors.New("only admin can access")
		}
		return uint(adminid), nil
	}
	return 0, errors.New("token invalid")
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
