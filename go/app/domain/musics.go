package domain

import (
	"time"

	"github.com/labstack/echo"
)

type Musics struct {
	Id          int64     `json:"id,omitempty" gorm:"AUTO_INCREMENT"`
	Title       string    `json:"title,omitempty" form:"title" validate:"required"`
	Description string    `json:"description,omitempty" form:"description" validate:"required"`
	TrackLength int64     `json:"track_length,omitempty" form:"track_length" validate:"required"`
	Album 			string    `json:"album,omitempty" form:"album"`
	FilePath 		string 	  `json:"file_path,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
}

type MusicFilePath struct {
	FilePath 		string 	  `json:"file_path,omitempty"`
}

type MusicRepository interface {
	GetAll(ctx echo.Context) (data []Musics, err error)
	GetById(ctx echo.Context, id int64) (music Musics, err error)
	GetAllByTitle(ctx echo.Context, title string) (music []Musics, err error)
	Post(ctx echo.Context, music Musics) (err error)
	Patch(ctx echo.Context, music Musics) (err error)
	Delete(ctx echo.Context, music Musics) (err error)
}

type MusicUseCase interface {
	GetAll(ctx echo.Context) (data []Musics, err error)
	GetById(ctx echo.Context, id int64) (music Musics, err error)
	SearchByTitle(ctx echo.Context, title string) (musics []Musics, err error)
	Post(ctx echo.Context, music Musics) (err error)
	Patch(ctx echo.Context, music Musics) (err error)
	Delete(ctx echo.Context, music Musics) (err error)
	GetMusicFile(ctx echo.Context, filePath string) (err error)
}