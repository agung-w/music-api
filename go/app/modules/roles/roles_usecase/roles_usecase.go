package roles_usecase

import (
	"main/app/domain"
	"main/app/global"

	"github.com/labstack/echo"
)

type roleUsecase struct {
}

func (r roleUsecase) GetByName(ctx echo.Context, name string) (role domain.Roles, err error) {
	role, err = global.RoleRepo.GetByName(ctx, name)
	if err != nil {
		return
	}
	return
}

func (r roleUsecase) Get(ctx echo.Context) (data []domain.Roles, err error) {
	data, err = global.RoleRepo.Get(ctx)
	if err != nil {
		return
	}
	return
}

func (r roleUsecase) GetById(ctx echo.Context, id int64) (role domain.Roles, err error) {
	role, err = global.RoleRepo.GetById(ctx, id)
	if err != nil {
		// util.LoggerI(ctx, err.Error())
		return 
	}
	return
}

func (r roleUsecase) Post(ctx echo.Context, role domain.Roles) (err error) {
	err = global.RoleRepo.Post(ctx, role)
	if err != nil {
		return
	}
	return
}

func (r roleUsecase) Patch(ctx echo.Context, role domain.Roles) (err error) {

	//Check if exist
	_, err = global.RoleRepo.GetById(ctx, role.Id)
	if err != nil {
		return
	}

	err = global.RoleRepo.Patch(ctx, role)
	if err != nil {
		return
	}
	return
}

func (r roleUsecase) Delete(ctx echo.Context, role domain.Roles) (err error) {
	err = global.RoleRepo.Delete(ctx, role)
	if err != nil {
		return
	}
	return
}

func New() domain.RoleUsecase {
	return &roleUsecase{}
}
