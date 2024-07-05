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
	MusicRepo domain.MusicRepository
	RoleRepo  domain.RoleRepository
	UserRepo  domain.UserRepository
)

var (
	MusicUseCase domain.MusicUseCase
	RoleUsecase  domain.RoleUsecase
	UserUsecase  domain.UserUseCase
)

const(
	JwtSecret = "secret-123"
)
	