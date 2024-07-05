package global

import (
	"main/app/domain"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)
var (
	DbConn      *gorm.DB
	Echo        *echo.Echo
	Validate    *validator.Validate
)

var (
	RoleRepo domain.RoleRepository
	UserRepo domain.UserRepository
)

var (
	RoleUsecase domain.RoleUsecase
	UserUsecase domain.UserUseCase
)

const(
	JwtSecret = "secret-123"
)
	