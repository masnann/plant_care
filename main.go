package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/masnann/plant_care/config"
	hAuth "github.com/masnann/plant_care/features/auth/handler"
	rAuth "github.com/masnann/plant_care/features/auth/repository"
	sAuth "github.com/masnann/plant_care/features/auth/service"
	hPlant "github.com/masnann/plant_care/features/plant/handler"
	rPlant "github.com/masnann/plant_care/features/plant/repository"
	sPlant "github.com/masnann/plant_care/features/plant/service"
	hUser "github.com/masnann/plant_care/features/user/handler"
	rUser "github.com/masnann/plant_care/features/user/repository"
	sUser "github.com/masnann/plant_care/features/user/service"
	"github.com/masnann/plant_care/middlewares"
	"github.com/masnann/plant_care/routes"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/database"
)

func main() {
	e := echo.New()
	var initConfig = config.InitConfig()

	db := database.InitDatabase(*initConfig)
	database.Migrate(db)

	jwtService := utils.NewJWT(initConfig.Secret)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo)
	userHandler := hUser.NewUserHandler(userService, jwtService)

	plantRepo := rPlant.NewPlantRepository(db)
	plantService := sPlant.NewPlantService(plantRepo)
	plantHandler := hPlant.NewPlantHandler(plantService, jwtService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, jwtService)
	authHandler := hAuth.NewAuthHandler(authService, userService, jwtService, *initConfig)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middlewares.ConfigureLogging())

	routes.RouteUser(e, userHandler, *initConfig)
	routes.RouteAuth(e, authHandler, *initConfig)
	routes.RoutePlant(e, plantHandler, jwtService, userService)
	e.Logger.Fatalf(e.Start(fmt.Sprintf(":%d", initConfig.ServerPort)).Error())
}
