package users_usecase

import (
	"main/app/domain"
	"main/app/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
}

// SetAsAdministrator implements domain.UserUseCase.
func (u *userUsecase) SetAsAdministrator(ctx echo.Context, email string) (err error) {
	user, err := global.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return
	}

	role, err := global.RoleRepo.GetByName(ctx, domain.ROLES_ADMINISTRATOR)
	if err != nil {
		return
	}

	user.RolesId=role.Id
	return global.UserRepo.Patch(ctx,&user)
}

// ResetPassword implements domain.UserUseCase.
func (u *userUsecase) ResetPassword(ctx echo.Context, email string) (err error) {
	panic("unimplemented")
}

func (*userUsecase) Login(ctx echo.Context, email, password string) (accessToken string, err error) {
	user, err := global.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return
	}

	role, err := global.RoleRepo.GetById(ctx, user.RolesId)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["role"] = role.Name
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	accessToken, err = token.SignedString([]byte(global.JwtSecret))
	if err != nil {
		return
	}

	return
}

func (*userUsecase) RegisterUser(ctx echo.Context, user domain.Users) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = global.UserRepo.GetByEmail(ctx, user.Email)
	if err == nil {
		return domain.ErrConflict
	}

	role, err := global.RoleRepo.GetByName(ctx, domain.ROLES_USER)
	if err != nil {
		return
	}

	user.CreatedTime = time.Now()
	user.Password = string(hashedPassword)
	user.RolesId = role.Id
	return global.UserRepo.Post(ctx, &user)
}

func (*userUsecase) Get(ctx echo.Context) (users []domain.Users, err error) {
	return global.UserRepo.Get(ctx)
}

func New() domain.UserUseCase {
	return &userUsecase{}
}
