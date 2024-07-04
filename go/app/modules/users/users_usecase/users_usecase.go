package users_usecase

import (
	"main/app/domain"
	"main/app/global"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)
type userUsecase struct {
}
func (*userUsecase) Login(ctx echo.Context, email, password string) (accessToken string, refreshToken string, err error){
	return
}
func (*userUsecase) RegisterUser(ctx echo.Context, user domain.Users) (err error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.CreatedTime = time.Now()
	user.Password = string(hashedPassword)
	return global.UserRepo.Post(ctx, &user)
}

func New() domain.UserUseCase{
	return &userUsecase{}
}
