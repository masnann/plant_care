package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/config"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/plant"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/middlewares"
	"github.com/masnann/plant_care/utils"
)

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface, cfg config.Config) {
	e.POST("/auth/register", h.Register())
	e.POST("/auth/login", h.Login())
}

func RouteUser(e *echo.Echo, h user.HandlerUserInterface, cfg config.Config) {
	var users = e.Group("/users")
	users.Use(echojwt.JWT([]byte(cfg.Secret)))

	users.GET("", h.GetAllUsers())
	users.GET("/by-email", h.GetUserByEmail())
}

func RoutePlant(e *echo.Echo, p plant.HandlerPlantInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.GET("/plants", p.GetPaginationPlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.GET("/plants/by-name", p.SearchPlantsByName(), middlewares.AuthMiddleware(jwtService, userService))
	e.GET("/plants/by-type", p.SearchPlantsByType(), middlewares.AuthMiddleware(jwtService, userService))
	e.POST("/plants", p.InsertPlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.PUT("/plants/updates/:id", p.UpdatePlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.DELETE("/plants/delete/:id", p.DeletePlants(), middlewares.AuthMiddleware(jwtService, userService))
}
