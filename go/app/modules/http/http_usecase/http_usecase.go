package http_usecase

import (
	"main/app/global"

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


