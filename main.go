package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/masnann/plant_care/config"
	hAuth "github.com/masnann/plant_care/features/auth/handler"
	rAuth "github.com/masnann/plant_care/features/auth/repository"
	sAuth "github.com/masnann/plant_care/features/auth/service"
	hUser "github.com/masnann/plant_care/features/user/handler"
	rUser "github.com/masnann/plant_care/features/user/repository"
	sUser "github.com/masnann/plant_care/features/user/service"
	"github.com/masnann/plant_care/routes"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/database"
)

func main() {
	e := echo.New()
	var initConfig = config.InitConfig()

	db := database.InitDatabase(*initConfig)
	database.Migrate(db)

	jwtInterface := utils.NewJWT(initConfig.Secret, initConfig.RefreshSecret)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo, jwtInterface)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, jwtInterface)
	authHandler := hAuth.NewAuthHandler(authService, userService, jwtInterface)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userHandler, *initConfig)
	routes.RouteAuth(e, authHandler, *initConfig)
	e.Logger.Fatalf(e.Start(fmt.Sprintf(":%d", initConfig.ServerPort)).Error())
}
