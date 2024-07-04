package users_repository

import (
	"main/app/domain"
	"main/app/global"

	"github.com/labstack/echo"
)
type gormUserRepo struct {
}

func (g gormUserRepo) Get(ctx echo.Context) (users []domain.Users, err error) {
	err = global.DbConn.Omit("id", "password", "2fa_on", "otp_secret", "otp_auth_url", "recovery_code").Find(&users).Error
	return
}

func (g gormUserRepo) GetById(ctx echo.Context, id int64) (user domain.Users, err error) {
	err = global.DbConn.Where("id=?", id).
		Omit("password", "2fa_on", "otp_secret", "otp_auth_url", "recovery_code").
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
