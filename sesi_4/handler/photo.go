package handler

import (
	"net/http"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/models"
	"sesi_4_final_project/photo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService photo.Service
}

func NewPhotoHandler(photoService photo.Service) *photoHandler {
	return &photoHandler{photoService}
}

// GetAllPhoto godoc
// @Summary      get all photo by user id
// @Description  get all photo by user id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/photo [get]
// @Security Bearer
func (h *photoHandler) GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	photos, err := h.photoService.GetAll(currentUser)

	if err != nil {
		response := helpers.JSONResult400{
			Message: "Get All Photos Failed",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Get All Photos By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    photos,
	}
	c.JSON(http.StatusOK, response)
}

// CreatePhoto godoc
// @Summary      create socialmedia by user id
// @Description  create socialmedia by user id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param		 createPhoto	body photo.CreatePhotoInput	true "Create Photo"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/photo [post]
// @Security 	 Bearer
func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Create photo failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhoto, err := h.photoService.CreatePhoto(input, currentUser)

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
		Message: "Success Create Photo By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    newPhoto,
	}
	c.JSON(http.StatusOK, response)
}

// GetOnePhoto godoc
// @Summary      get one photo by user id
// @Description  get one photo by user id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param 		 photoID  path int true "photoID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/photo/{photoID} [get]
// @Security Bearer
func (h *photoHandler) GetOnePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("photoID"))
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

	photo, err := h.photoService.GetOne(uint(photoID))

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
		Message: "Get One Photo By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    photo,
	}
	c.JSON(http.StatusOK, response)
}

// DeletePhoto godoc
// @Summary      delete photo by user id
// @Description  delete photo by user id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param 		 photoID  path int true "photoID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/photo/{photoID} [delete]
// @Security Bearer
func (h *photoHandler) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	photoID, _ := strconv.Atoi(c.Param("photoID"))

	err := h.photoService.DeletePhoto(uint(photoID), currentUser)

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
		Message: "Succesfully Delete photo By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

// UpdatePhoto godoc
// @Summary      update photo by user id
// @Description  update photo by user id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param 		 photoID  path int true "photoID"
// @Param		 updatePhoto	body photo.CreatePhotoInput	true "Update Photo"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/photo/{photoID} [put]
// @Security Bearer
func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var input photo.CreatePhotoInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	photoID, _ := strconv.Atoi(c.Param("photoID"))
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Cannot update photo failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatePhoto, err := h.photoService.UpdatePhoto(uint(photoID), input, currentUser)

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
		Message: "Succesfully Update photo By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    updatePhoto,
	}
	c.JSON(http.StatusOK, response)
}
