package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtService utils.JWTInterface, userService user.ServiceUserInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing or Invalid Bearer Token")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid Token "+err.Error())
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid or Expired Token "+err.Error())
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid User ID "+err.Error())
			}

			userID := uint64(userIDFloat)

			users, err := userService.GetUserById(userID)
			if err != nil {
				return response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User Not Found "+err.Error())
			}

			c.Set("CurrentUser", users)

			return next(c)
		}
	}
}
