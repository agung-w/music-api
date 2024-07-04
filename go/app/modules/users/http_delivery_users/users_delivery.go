package http_delivery_users

import (
	"main/app/domain"
	"main/app/global"
	"main/app/modules/http/http_usecase"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUsecase domain.UserUseCase
}

func (h UserHandler) Login(c echo.Context) error {
	data := new(domain.Users)
	if err := c.Bind(data); err != nil {
		return err
	}

	accessToken, err := global.UserUsecase.Login(c, data.Email, data.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			// "rc":  domain.RC_02_INVALID_AUTHORIZATION,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			// "rc":            domain.RC_00_OK,
			"access_token":  accessToken,
		})
	}
}


func (h UserHandler) CreateUser(c echo.Context) error {
	data := new(domain.Users)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := global.Validate.Struct(data); err != nil {
		// util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  "domain.RC_01_INVALID_PAYLOAD",
			"msg": err.Error(),
		})
	}

	err := global.UserUsecase.RegisterUser(c, *data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  "domain.RC_03_INTERNAL_ERROR",
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			// "rc": "domain.RC_00_OK",
		})
	}
}

func (h UserHandler) ListUsers(c echo.Context) error {
	users, err := global.UserUsecase.Get(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "domain.ErrBadParamInput.Error()")
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			// "rc":    domain.RC_00_OK,
			"users": users,
		})
	}
}


func HttpUserHandler() {
	handler := &UserHandler{}
	global.Echo.POST("/users/register", handler.CreateUser)
	global.Echo.POST("/users/login", handler.Login)
	global.Echo.GET("/users", handler.ListUsers, http_usecase.IsLoggedIn)
}
