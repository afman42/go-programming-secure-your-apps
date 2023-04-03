package middlewares

import (
	"net/http"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/models"
	"sesi_4_final_project/socialmedia"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthorizationSocialMediaByUserID(socialMediaService socialmedia.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
		if err != nil {
			response := helpers.JSONResult400{
				Code:    http.StatusBadRequest,
				Message: "Invalid Parameter",
				Data:    nil,
				Status:  "error",
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData := c.MustGet("userData").(models.User)
		userID := uint(userData.ID)

		socialMedia, err := socialMediaService.GetOne(uint(socialMediaID))

		if err != nil {
			response := helpers.JSONResult404{
				Code:    http.StatusNotFound,
				Message: "Data does'nt exist",
				Data:    nil,
				Status:  "error",
			}
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		if socialMedia.UserID != userID {
			response := helpers.JSONResult401{
				Status:  "error",
				Message: "you are not allowed to access this data",
				Code:    http.StatusUnauthorized,
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Next()
	}
}
