package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/assistant"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/guide"
	"github.com/masnann/plant_care/features/note"
	"github.com/masnann/plant_care/features/notification"
	"github.com/masnann/plant_care/features/plant"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/middlewares"
	"github.com/masnann/plant_care/utils"
)

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface) {
	e.POST("/auth/register", h.Register())
	e.POST("/auth/login", h.Login())
	e.GET("/verify-email", h.VerifyEmail())
}

func RoutePlant(e *echo.Echo, p plant.HandlerPlantInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.GET("/plants", p.GetPaginationPlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.GET("/plants/by-name", p.SearchPlantsByName(), middlewares.AuthMiddleware(jwtService, userService))
	e.GET("/plants/by-type", p.SearchPlantsByType(), middlewares.AuthMiddleware(jwtService, userService))
	e.POST("/plants", p.InsertPlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.PUT("/plants/updates/:id", p.UpdatePlants(), middlewares.AuthMiddleware(jwtService, userService))
	e.DELETE("/plants/delete/:id", p.DeletePlants(), middlewares.AuthMiddleware(jwtService, userService))
}

func RouteAssistant(e *echo.Echo, a assistant.HandlerAssistantInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.POST("/assistant", a.Assistant(), middlewares.AuthMiddleware(jwtService, userService))
}

func RouteGuide(e *echo.Echo, g guide.HandlerGuideInterface) {
	e.GET("/guide", g.GetGuidesWithPagination())
	e.GET("/guide/:id", g.GetGuidesById())
	e.GET("/guide/search", g.SearchGuideByName())
}

func RouteNote(e *echo.Echo, n note.HandlerNoteInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.POST("/notes", n.InsertNotes(), middlewares.AuthMiddleware(jwtService, userService))
	e.GET("/notes", n.GetNotesWithPagination(), middlewares.AuthMiddleware(jwtService, userService))
	e.PUT("/notes/updates/:id", n.UpdateNotes(), middlewares.AuthMiddleware(jwtService, userService))
	e.DELETE("/notes/delete/:id", n.DeleteNotes(), middlewares.AuthMiddleware(jwtService, userService))

	e.POST("/notes/photo", n.InsertNotePhoto(), middlewares.AuthMiddleware(jwtService, userService))
	e.PUT("/notes/photos/updates/:id", n.UpdateNotesPhotos(), middlewares.AuthMiddleware(jwtService, userService))
	e.DELETE("/notes/photos/delete/:id", n.DeleteNotesPhotos(), middlewares.AuthMiddleware(jwtService, userService))
}

func RouteNotify(e *echo.Echo, n notification.HandlerNotificationInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.GET("/notifications", n.GetPaginationNotifications(), middlewares.AuthMiddleware(jwtService, userService))
}

func RouteUsers(e *echo.Echo, n user.HandlerUserInterface, jwtService utils.JWTInterface, userService user.ServiceUserInterface) {
	e.POST("/change-pass", n.UpdatePassword(), middlewares.AuthMiddleware(jwtService, userService))
}
