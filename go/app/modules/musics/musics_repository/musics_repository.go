package musics_repository

import (
	"main/app/domain"
	"main/app/global"

	"github.com/labstack/echo"
)

type gormMusicRepo struct {
}

func (g gormMusicRepo) Delete(ctx echo.Context, music domain.Musics) (err error) {
	err = global.DbConn.Delete(&music).Error
	return
}

func (g gormMusicRepo) GetAll(ctx echo.Context) (data []domain.Musics, err error) {
	err = global.DbConn.Omit("description", "file_path").Find(&data).Error
	return
}

func (g gormMusicRepo) GetById(ctx echo.Context, id int64) (music domain.Musics, err error) {
	err = global.DbConn.Where("id=?", id).First(&music).Error
	return
}

func (g gormMusicRepo) GetAllByTitle(ctx echo.Context, title string) (musics []domain.Musics, err error) {
	err = global.DbConn.Omit("description", "file_path").Where("title LIKE ?", "%"+title+"%").Find(&musics).Error
	return
}

func (g gormMusicRepo) Patch(ctx echo.Context, music domain.Musics) (err error) {
	err = global.DbConn.Updates(music).Error
	return
}

func (g gormMusicRepo) Post(ctx echo.Context, music domain.Musics) (err error) {
	err = global.DbConn.Create(&music).Error
	return
}

func New() domain.MusicRepository {
	return gormMusicRepo{}
}
