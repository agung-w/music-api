package roles_repository

import (
	"main/app/domain"
	"main/app/global"

	"github.com/labstack/echo"
)

type gormRoleRepo struct {
}

func (g gormRoleRepo) GetByName(ctx echo.Context, name string) (role domain.Roles, err error) {
	err = global.DbConn.Where("name=?", name).First(&role).Error
	return
}

func (g gormRoleRepo) Get(ctx echo.Context) (data []domain.Roles, err error) {
	err = global.DbConn.Find(&data).Error
	return
}

func (g gormRoleRepo) GetById(ctx echo.Context, id int64) (role domain.Roles, err error) {
	err = global.DbConn.Where("id=?", id).First(&role).Error
	return
}

func (g gormRoleRepo) Post(ctx echo.Context, role domain.Roles) (err error) {
	err = global.DbConn.Create(&role).Error
	return
}

func (g gormRoleRepo) Patch(ctx echo.Context, role domain.Roles) (err error) {
	err = global.DbConn.Updates(&role).Error
	return
}

func (g gormRoleRepo) Delete(ctx echo.Context, role domain.Roles) (err error) {
	err = global.DbConn.Delete(&role).Error
	return
}

func New() domain.RoleRepository {
	return &gormRoleRepo{}
}
