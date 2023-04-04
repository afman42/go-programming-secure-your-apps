package handler

import (
	"fmt"
	"net/http"
	"sesi_4_final_project/comments"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentsHandler struct {
	commentsService comments.Service
}

func NewCommentsHandler(commentsService comments.Service) *commentsHandler {
	return &commentsHandler{commentsService}
}

// GetAllComment godoc
// @Summary      get all comments by user id
// @Description  get all comments by user id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 photoID  path int true "photoID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/comment/{photoID} [get]
// @Security Bearer
func (h *commentsHandler) GetAllComment(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		response := helpers.JSONResult400{
			Message: "Invalid Parameter",
			Code:    http.StatusBadRequest,
			Data:    nil,
			Status:  "error",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	comments, err := h.commentsService.GetAll(uint(photoID))

	if err != nil {
		response := helpers.JSONResult400{
			Message: "Get All Comments Failed",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Get All Comments By Photo",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    comments,
	}
	c.JSON(http.StatusOK, response)
}

// CreateComment godoc
// @Summary      create comment by user id and photo id
// @Description  create comment by user id and photo id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 photoID  path int true "photoID"
// @Param		 createComment	body comments.CreateCommentInput	true "Create Social Media"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/comment/{photoID} [post]
// @Security Bearer
func (h *commentsHandler) CreateComment(c *gin.Context) {
	var input comments.CreateCommentInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID

	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		response := helpers.JSONResult400{
			Message: "Invalid Parameter",
			Code:    http.StatusBadRequest,
			Data:    nil,
			Status:  "error",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Create comment failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newComment, err := h.commentsService.CreateComment(input, currentUser, uint(photoID))

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
		Message: "Success Create Comment By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    newComment,
	}
	c.JSON(http.StatusOK, response)
}

// GetOneComment godoc
// @Summary      get one comment by user id
// @Description  get one comment by user id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 photoID  	path int true "photoID"
// @Param 		 commentID  path int true "commentID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/comment/{photoID}/{commentID} [get]
// @Security 	 Bearer
func (h *commentsHandler) GetOneComment(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		response := helpers.JSONResult400{
			Message: fmt.Sprintf("Invalid Parameter photo id: %d", photoID),
			Status:  "error",
			Data:    nil,
			Code:    http.StatusBadRequest,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	commentID, err := strconv.Atoi(c.Param("commentID"))
	if err != nil {
		response := helpers.JSONResult400{
			Message: fmt.Sprintf("Invalid Parameter comment id: %d", commentID),
			Status:  "error",
			Data:    nil,
			Code:    http.StatusBadRequest,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	comment, err := h.commentsService.GetOne(uint(commentID), uint(photoID))

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
		Message: "Get One Comment By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    comment,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteComment godoc
// @Summary      delete comment by user id and photo id
// @Description  delete comment by user id and photo id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 photoID  	path int true "photoID"
// @Param 		 commentID  path int true "commentID"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Router       /api/comment/{photoID}/{commentID} [delete]
// @Security 	 Bearer
func (h *commentsHandler) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	photoID, _ := strconv.Atoi(c.Param("photoID"))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	err := h.commentsService.DeleteComment(uint(commentID), currentUser, uint(photoID))

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
		Message: "Succesfully Delete Comment By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

// UpdateComment godoc
// @Summary      update comment by user id
// @Description  update comment by user id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 photoID  	path int true "photoID"
// @Param 		 commentID  path int true "commentID"
// @Param		 updateComment	body comments.CreateCommentInput	true "Update Comment"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 400  {object}  helpers.JSONResult400
// @Failure		 404  {object}  helpers.JSONResult404
// @Failure		 401  {object}  helpers.JSONResult401
// @Router       /api/comment/{photoID}/{commentID} [put]
// @Security 	 Bearer
func (h *commentsHandler) UpdateComment(c *gin.Context) {
	var input comments.CreateCommentInput
	userData := c.MustGet("userData").(models.User)
	currentUser := userData.ID
	photoID, _ := strconv.Atoi(c.Param("photoID"))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Cannot update comment failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateComment, err := h.commentsService.UpdateComment(uint(commentID), input, currentUser, uint(photoID))

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
		Message: "Succesfully Update Comment By User",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    updateComment,
	}
	c.JSON(http.StatusOK, response)
}
