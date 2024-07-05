package main

import (
	"main/app/domain"
	"main/app/global"
	"main/app/modules/roles/roles_repository"
	"main/app/modules/roles/roles_usecase"
	"main/app/modules/users/http_delivery_users"
	"main/app/modules/users/users_repository"
	"main/app/modules/users/users_usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	db, err := gorm.Open(sqlite.Open("music-api.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
	db.AutoMigrate(&domain.Users{}, &domain.Roles{})
	global.DbConn=db

	global.Echo = echo.New()
	global.Echo.Use(middleware.CORS())


	global.Validate = validator.New(validator.WithRequiredStructEnabled())
	registerRepo()

	//Register Usecase
	registerUsecase()

	//Register Handler
	registerHTTPHandler()

	global.Echo.Logger.Fatal(global.Echo.Start(":1323"))
	
}

// func save(c echo.Context) error {
// 	header := c.Request().Header
// 	name := header.Get("name")
// 	email := header.Get("email")
// 	password := header.Get("password")
// 	newUser := domain.User{
//     Name:  name,
//     Email: email,
// 		Password: password,
//   }
// 	db.Create(newUser)

// 	item := domain.User{}
// 	db.First(&item, "email = ?", email)
// 	fmt.Println(item)// find product with integer primary key
	

// 	return c.JSON(http.StatusOK,  newUser)
// }

func registerRepo() {
	global.RoleRepo = roles_repository.New()
	global.UserRepo = users_repository.New()
}

func registerUsecase() {
	global.RoleUsecase = roles_usecase.New()
	global.UserUsecase = users_usecase.New()
}

func registerHTTPHandler() {
	http_delivery_users.HttpUserHandler()
}