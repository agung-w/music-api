package domain

import "github.com/labstack/echo"

const (
	ROLES_ADMINISTRATOR   = "administrator"
	ROLES_USER            = "user"
	ROLE_ID_ADMINISTRATOR = 1
	ROLE_ID_USER          = 2
)

type Roles struct {
	Id   int64  `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name" validate:"required"`
}


type RoleRepository interface {
	Get(ctx echo.Context) (data []Roles, err error)
	GetById(ctx echo.Context, id int64) (role Roles, err error)
	GetByName(ctx echo.Context, name string) (role Roles, err error)
	Post(ctx echo.Context, role Roles) (err error)
	Patch(ctx echo.Context, role Roles) (err error)
	Delete(ctx echo.Context, role Roles) (err error)
}

type RoleUsecase interface {
	Get(ctx echo.Context) (data []Roles, err error)
	GetById(ctx echo.Context, id int64) (role Roles, err error)
	GetByName(ctx echo.Context, name string) (role Roles, err error)
	Post(ctx echo.Context, role Roles) (err error)
	Patch(ctx echo.Context, role Roles) (err error)
	Delete(ctx echo.Context, role Roles) (err error)
}


