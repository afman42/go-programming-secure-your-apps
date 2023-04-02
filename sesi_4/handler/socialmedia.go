package handler

import (
	"net/http"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/models"
	"sesi_4_final_project/socialmedia"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService socialmedia.Service
}

func NewSocialMediaHandler(socialMediaService socialmedia.Service) *socialMediaHandler {
	return &socialMediaHandler{socialMediaService}
}

// GetAllSocialMedia godoc
// @Summary      get all socialmedia by user id
// @Description  get all socialmedia by user id
// @Tags         socialmedia
// @Accept       json
// @Produce      json
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/social_media [get]
// @Security Bearer
func (h *socialMediaHandler) GetAllSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	socialmedias, err := h.socialMediaService.GetAll(currentUser)

	if err != nil {
		response := helpers.JSONResult400{
			Message: "Get All Social Media Failed",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Get All Social Media By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    socialmedias,
	}
	c.JSON(http.StatusOK, response)
}

// // UserLogin godoc
// // @Summary      user login
// // @Description  get token for login
// // @Tags         authentication
// // @Accept       json
// // @Produce      json
// // @Param		 login	body	user.LoginInput	true "Auth Login"
// // @Success      200  {object} 	helpers.JSONResult200
// // @Failure		 422  {object}  helpers.JSONResult422
// // @Failure		 400  {object}  helpers.JSONResult400
// // @Router       /api/login [post]
// func (h *userHandler) Login(c *gin.Context) {
// 	var input user.LoginInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		response := helpers.JSONResult422{
// 			Message: "Login failed",
// 			Code:    http.StatusUnprocessableEntity,
// 			Status:  "error",
// 			Errors:  helpers.FormatValidationError(err),
// 		}
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	loggedinUser, err := h.userService.Login(input)

// 	if err != nil {
// 		response := helpers.JSONResult422{
// 			Message: "Login failed",
// 			Code:    http.StatusUnprocessableEntity,
// 			Status:  "error",
// 			Errors:  helpers.FormatValidationError(err),
// 		}
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Email)
// 	if err != nil {
// 		response := helpers.JSONResult400{
// 			Message: "Login failed",
// 			Code:    http.StatusBadRequest,
// 			Status:  "error",
// 			Data:    nil,
// 		}
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := helpers.JSONResult200{
// 		Message: "Successfuly logged",
// 		Code:    http.StatusOK,
// 		Status:  "success",
// 		Data:    token,
// 	}
// 	c.JSON(http.StatusOK, response)
// }
