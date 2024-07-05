package musics_usecase

import (
	"errors"
	"io"
	"main/app/domain"
	"main/app/global"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type musicUseCase struct {
}

// GetMusicFile implements domain.MusicUseCase.
func (m musicUseCase) GetMusicFile(ctx echo.Context, filePath string) (err error) {
	return ctx.File(filePath)
}

// Delete implements domain.MusicUseCase.
func (m musicUseCase) Delete(ctx echo.Context, music domain.Musics) (err error) {
	err = global.MusicRepo.Delete(ctx, music)
	if err != nil {
		return
	}
	return
}

// GetAll implements domain.MusicUseCase.
func (m musicUseCase) GetAll(ctx echo.Context) (data []domain.Musics, err error) {
	data, err = global.MusicRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

// GetById implements domain.MusicUseCase.
func (m musicUseCase) GetById(ctx echo.Context, id int64) (music domain.Musics, err error) {
	music, err = global.MusicRepo.GetById(ctx, id)
	if err != nil {
		return
	}
	return
}

// Patch implements domain.MusicUseCase.
func (m musicUseCase) Patch(ctx echo.Context, music domain.Musics) (err error) {
	_, err = global.MusicRepo.GetById(ctx, music.Id)
	if err != nil {
		return
	}
	file, handler, err := ctx.Request().FormFile("music_file")
	if err != nil {
		err = global.MusicRepo.Patch(ctx, music)
		if err != nil {
			return
		}
		return
	}
	defer file.Close()

	now := time.Now()
	newFileName := strconv.Itoa(int(now.UnixNano())) + "_" + handler.Filename
	dstPath := filepath.Join("musics", newFileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return err
	}

	music.FilePath = dstPath
	err = global.MusicRepo.Patch(ctx, music)
	if err != nil {
		return
	}
	return
}

// Post implements domain.MusicUseCase.
func (m musicUseCase) Post(ctx echo.Context, music domain.Musics) (err error) {
	file, handler, err := ctx.Request().FormFile("music_file")
	if err != nil {
		return errors.New("failed to retrieved the file")
	}
	defer file.Close()

	now := time.Now()
	newFileName := strconv.Itoa(int(now.UnixNano())) + "_" + handler.Filename
	dstPath := filepath.Join("musics", newFileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return err
	}

	music.CreatedTime = now
	music.FilePath = dstPath
	err = global.MusicRepo.Post(ctx, music)
	if err != nil {
		return
	}
	return
}

// SearchByTitle implements domain.MusicUseCase.
func (m musicUseCase) SearchByTitle(ctx echo.Context, title string) (musics []domain.Musics, err error) {
	musics, err = global.MusicRepo.GetAllByTitle(ctx, title)
	if err != nil {
		return
	}
	return
}

func New() domain.MusicUseCase {
	return musicUseCase{}
}
