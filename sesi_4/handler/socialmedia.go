package handler

import (
	"net/http"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/models"
	"sesi_4_final_project/socialmedia"
	"strconv"

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

// CreateSocialMedia godoc
// @Summary      create socialmedia by user id
// @Description  create socialmedia by user id
// @Tags         socialmedia
// @Accept       json
// @Produce      json
// @Param		 createSocialMedia	body socialmedia.CreateSocialMediaInput	true "Create Social Media"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/social_media [post]
// @Security Bearer
func (h *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var input socialmedia.CreateSocialMediaInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Create social media failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSocialMedia, err := h.socialMediaService.CreateSocialMedia(input, currentUser)

	if err != nil {
		response := helpers.JSONResult400{
			Message: "Cannot Create Social Media",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Success Create Social Media By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    newSocialMedia,
	}
	c.JSON(http.StatusOK, response)
}

// GetOneSocialMedia godoc
// @Summary      get one socialmedia by user id
// @Description  get one socialmedia by user id
// @Tags         socialmedia
// @Accept       json
// @Produce      json
// @Param 		 socialMediaID  path int true "socialMediaID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/social_media/{socialMediaID} [get]
// @Security Bearer
func (h *socialMediaHandler) GetOneSocialMedia(c *gin.Context) {
	socialMediaID, err := strconv.Atoi(c.Param("socialMediaID"))
	if err != nil {
		response := helpers.JSONResult400{
			Message: "Invalid Parameter",
			Status:  "error",
			Data:    nil,
			Code:    http.StatusBadRequest,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	socialMedia, err := h.socialMediaService.GetOne(uint(socialMediaID))

	if err != nil {
		response := helpers.JSONResult404{
			Message: err.Error(),
			Code:    http.StatusNotFound,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Get One Social Media By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    socialMedia,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteSocialMedia godoc
// @Summary      delete socialmedia by user id
// @Description  delete socialmedia by user id
// @Tags         socialmedia
// @Accept       json
// @Produce      json
// @Param 		 socialMediaID  path int true "socialMediaID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/social_media/{socialMediaID} [delete]
// @Security Bearer
func (h *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))

	err := h.socialMediaService.DeleteSocialMedia(uint(socialMediaID), currentUser)

	if err != nil {
		response := helpers.JSONResult404{
			Message: err.Error(),
			Code:    http.StatusNotFound,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Succesfully Delete Social Media By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateSocialMedia godoc
// @Summary      update socialmedia by user id
// @Description  update socialmedia by user id
// @Tags         socialmedia
// @Accept       json
// @Produce      json
// @Param 		 socialMediaID  path int true "socialMediaID"
// @Param		 updateSocialMedia	body socialmedia.CreateSocialMediaInput	true "Update Social Media"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/social_media/{socialMediaID} [put]
// @Security Bearer
func (h *socialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	var input socialmedia.CreateSocialMediaInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Cannot update social media failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateSocialMedia, err := h.socialMediaService.UpdateSocialMedia(uint(socialMediaID), input, currentUser)

	if err != nil {
		response := helpers.JSONResult400{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Succesfully Update Social Media By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    updateSocialMedia,
	}
	c.JSON(http.StatusOK, response)
}
