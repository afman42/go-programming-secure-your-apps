package middlewares

import (
	"net/http"
	"sesi_4_final_project/auth"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.JSONResult401{
				Message: "Unauthorized",
				Code:    http.StatusUnauthorized,
				Status:  "error",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helpers.JSONResult401{
				Message: "Unauthorized",
				Code:    http.StatusUnauthorized,
				Status:  "error",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helpers.JSONResult401{
				Message: "Unauthorized",
				Code:    http.StatusUnauthorized,
				Status:  "error",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := uint(claim["id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helpers.JSONResult401{
				Message: "Unauthorized",
				Code:    http.StatusUnauthorized,
				Status:  "error",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userData", user)
	}
}
