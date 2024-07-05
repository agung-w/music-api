package users_repository

import (
	"main/app/domain"
	"main/app/global"

	"github.com/labstack/echo"
)
type gormUserRepo struct {
}

func (g gormUserRepo) Get(ctx echo.Context) (users []domain.Users, err error) {
	err = global.DbConn.Omit("id", "password").Find(&users).Error
	return
}

func (g gormUserRepo) GetById(ctx echo.Context, id int64) (user domain.Users, err error) {
	err = global.DbConn.Where("id=?", id).
		Omit("password").
		First(&user).Error
	return
}

func (g gormUserRepo) GetByEmail(ctx echo.Context, email string) (user domain.Users, err error) {
	err = global.DbConn.Where("email=?", email).First(&user).Error
	return 
}

func (g gormUserRepo) Post(ctx echo.Context, user *domain.Users) (err error) {
	err = global.DbConn.Create(&user).Error
	return
}

func (g gormUserRepo) Patch(ctx echo.Context, user *domain.Users) (err error) {
	err = global.DbConn.Updates(&user).Error
	return
}

func New() domain.UserRepository {
	return &gormUserRepo{}
}
