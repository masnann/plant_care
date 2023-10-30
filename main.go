package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/masnann/plant_care/config"
	hAssistant "github.com/masnann/plant_care/features/assistant/handler"
	sAssistant "github.com/masnann/plant_care/features/assistant/service"
	hAuth "github.com/masnann/plant_care/features/auth/handler"
	rAuth "github.com/masnann/plant_care/features/auth/repository"
	sAuth "github.com/masnann/plant_care/features/auth/service"
	hGuide "github.com/masnann/plant_care/features/guide/handler"
	rGuide "github.com/masnann/plant_care/features/guide/repository"
	sGuide "github.com/masnann/plant_care/features/guide/service"
	hNote "github.com/masnann/plant_care/features/note/handler"
	rNote "github.com/masnann/plant_care/features/note/repository"
	sNote "github.com/masnann/plant_care/features/note/service"
	hNotify "github.com/masnann/plant_care/features/notification/handler"
	rNotify "github.com/masnann/plant_care/features/notification/repository"
	sNotify "github.com/masnann/plant_care/features/notification/service"
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

	notifyRepo := rNotify.NewNotificationRepository(db)
	notifyService := sNotify.NewNotificationService(notifyRepo)
	notifyHandler := hNotify.NewNotificationHandler(notifyService)

	plantRepo := rPlant.NewPlantRepository(db)
	plantService := sPlant.NewPlantService(plantRepo)
	plantHandler := hPlant.NewPlantHandler(plantService, jwtService, notifyService)

	assistantService := sAssistant.NewAssistantService()
	assistantHandler := hAssistant.NewAssistantHandler(assistantService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, jwtService)
	authHandler := hAuth.NewAuthHandler(authService, userService, jwtService, *initConfig)

	guideRepo := rGuide.NewGuideRepository(db)
	guideService := sGuide.NewGuideService(guideRepo)
	guideHandler := hGuide.NewGuideHandler(guideService)

	noteRepo := rNote.NewNoteRepository(db)
	noteService := sNote.NewNoteService(noteRepo)
	noteHandler := hNote.NewNoteHandler(noteService, plantService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middlewares.ConfigureLogging())

	routes.RouteUser(e, userHandler, *initConfig)
	routes.RouteAuth(e, authHandler, *initConfig)
	routes.RoutePlant(e, plantHandler, jwtService, userService)
	routes.RouteAssistant(e, assistantHandler, jwtService, userService)
	routes.RouteGuide(e, guideHandler)
	routes.RouteNote(e, noteHandler, jwtService, userService)
	routes.RouteNotify(e, notifyHandler, jwtService, userService)
	e.Logger.Fatalf(e.Start(fmt.Sprintf(":%d", initConfig.ServerPort)).Error())
}
