package http_usecase

import (
	"errors"
	"main/app/domain"
	"main/app/global"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
type GoMiddleware struct {
}

func (*GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(global.JwtSecret),
})

func IsAdministrator(next echo.HandlerFunc) echo.HandlerFunc {
	return checkRole(domain.ROLES_ADMINISTRATOR, next)
}

func checkRole(role string, next echo.HandlerFunc) func(c echo.Context) error {
	return func(c echo.Context) error {
		
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("JWT token missing or invalid")
		}

		claims := user.Claims.(jwt.MapClaims)
		claimRole, ok := claims["role"].(string)
		if !ok || role != claimRole {
			return echo.ErrUnauthorized
		}

		if ok {
			c.Set("role", claimRole)
		}

		val, ok := claims["name"].(string)
		if ok {
			c.Set("name", val)
		}

		val, ok = claims["email"].(string)
		if ok {
			c.Set("email", val)
		}

		return next(c)
	}
}

