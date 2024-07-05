package http_delivery_roles

import (
	"main/app/domain"
	"main/app/global"
	"main/app/modules/http/http_usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
}

func (h UserHandler) listRole(c echo.Context) error {
	users, err := global.RoleUsecase.Get(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			// "rc":    domain.RC_00_OK,
			"roles": users,
		})
	}
}

func (h UserHandler) createRole(c echo.Context) error {
	data := new(domain.Roles)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	if err := global.Validate.Struct(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	err := global.RoleUsecase.Post(c, *data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			// "rc": domain.RC_00_OK,
		})
	}
}

func (h UserHandler) patchRole(c echo.Context) error {
	data := new(domain.Roles)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	if err := global.Validate.Struct(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	err := global.RoleUsecase.Patch(c, *data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			// "rc": domain.RC_00_OK,
		})
	}
}

func (h UserHandler) deleteRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	data, err := global.RoleUsecase.GetById(c, int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	err = global.RoleUsecase.Delete(c, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			// "rc": domain.RC_00_OK,
		})
	}
}

func (h UserHandler) getRoleById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}

	role, err := global.RoleUsecase.GetById(c, int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			// "rc":   domain.RC_00_OK,
			"role": role,
		})
	}
}

func HttpRoleHandler() {
	handler := &UserHandler{}

	global.Echo.POST("/roles", handler.createRole, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.GET("/roles", handler.listRole, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.GET("/roles/:id", handler.getRoleById, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.PATCH("/roles", handler.patchRole, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.DELETE("/roles/:id", handler.deleteRole, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
}
