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
	UserRepo domain.UserRepository
)

var (
	UserUsecase domain.UserUseCase
)