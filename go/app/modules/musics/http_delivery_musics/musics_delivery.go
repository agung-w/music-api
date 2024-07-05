package http_delivery_musics

import (
	"main/app/domain"
	"main/app/global"
	"main/app/modules/http/http_usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type MusicHandler struct{
	MusicUseCase domain.MusicUseCase
}

func (h MusicHandler) getMusicFile(c echo.Context) error{
	data := new(domain.MusicFilePath)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	return global.MusicUseCase.GetMusicFile(c, data.FilePath)
}

func (h MusicHandler) getAllMusic(c echo.Context) error{
	data, err := global.MusicUseCase.GetAll(c)
	if err!=nil{
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}
	
	return c.JSON(http.StatusOK,  map[string]interface{}{
		// "rc":   domain.RC_00_OK,
		"musics": data,
	})
}

func (h MusicHandler) getById(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	data, err:= global.MusicUseCase.GetById(c,int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}else{
		return c.JSON(http.StatusOK,  map[string]interface{}{
			// "rc":   domain.RC_00_OK,
			"musics": data,
		})
	}
}
func (h MusicHandler) searchByTitle(c echo.Context) error{
	data, err:= global.MusicUseCase.SearchByTitle(c, c.Param("title"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			// "rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": domain.ErrorInvalidPayload,
		})
	}else{
		return c.JSON(http.StatusOK,  map[string]interface{}{
			// "rc":   domain.RC_00_OK,
			"musics": data,
		})
	}
}

func (h MusicHandler) createMusic(c echo.Context) error{
	data := new(domain.Musics)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}
	
	if err := global.Validate.Struct(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	err := global.MusicUseCase.Post(c, *data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
		})
	}
}

func (h MusicHandler) updateMusic(c echo.Context) error{
	data := new(domain.Musics)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	if err := global.Validate.Struct(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	err := global.MusicUseCase.Patch(c, *data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
		})
	}
}
func (h MusicHandler) DeleteMusic(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": domain.ErrorInvalidPayload,
		})
	}

	data, err := global.MusicUseCase.GetById(c, int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	} 

	err = global.MusicUseCase.Delete(c, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
		})
	}
}
func HttpUserHandler() {
	handler := &MusicHandler{}
	global.Echo.GET("/musics", handler.getAllMusic, http_usecase.IsLoggedIn)
	global.Echo.GET("/musics/:id", handler.getById, http_usecase.IsLoggedIn)
	global.Echo.GET("/search/musics/:title", handler.searchByTitle, http_usecase.IsLoggedIn)
	global.Echo.GET("/play/musics", handler.getMusicFile, http_usecase.IsLoggedIn)
	global.Echo.POST("/musics", handler.createMusic, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.PATCH("/musics", handler.updateMusic, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
	global.Echo.DELETE("/musics/:id", handler.DeleteMusic, http_usecase.IsLoggedIn, http_usecase.IsAdministrator)
}