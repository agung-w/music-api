package domain

import (
	"time"

	"github.com/labstack/echo"
)

type Users struct {
	Id          int64     `json:"id,omitempty" gorm:"AUTO_INCREMENT"`
	Email       string    `json:"email,omitempty" validate:"required,email"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Password    string    `json:"password,omitempty" validate:"required"`
	CreatedTime time.Time `json:"created_time,omitempty"`
}

type UserRepository interface{
	Get(ctx echo.Context) (users []Users,err error)
	GetById(ctx echo.Context, id int64) (user Users, err error)
	GetByEmail(ctx echo.Context, email string) (user Users, err error)
	Post(ctx echo.Context, user *Users) (err error)
	Patch(ctx echo.Context, user *Users) (err error)
}

type UserUseCase interface {
	Login(ctx echo.Context, email, password string) (accessToken string, refreshToken string, err error)
	RegisterUser(ctx echo.Context, user Users) (err error)
}
