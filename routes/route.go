package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/config"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/user"
)

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface, cfg config.Config) {
	e.POST("/auth/register", h.Register())
	e.POST("/auth/login", h.Login())
	e.POST("/refresh", h.RefreshJWT(), echojwt.JWT([]byte(cfg.RefreshSecret)))
}

func RouteUser(e *echo.Echo, h user.HandlerUserInterface, cfg config.Config) {
	var users = e.Group("/users")
	users.Use(echojwt.JWT([]byte(cfg.Secret)))

	users.GET("", h.GetAllUsers())
	users.GET("/by-email", h.GetUserByEmail())
}
