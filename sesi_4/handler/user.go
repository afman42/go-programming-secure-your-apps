package handler

import (
	"net/http"
	"sesi_4_final_project/auth"
	"sesi_4_final_project/helpers"
	"sesi_4_final_project/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// RegisterUser godoc
// @Summary      user register
// @Description  create user
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param		 register	body	user.RegisterUserInput	true "Auth Register"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Router       /api/register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Register account failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

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

	token, err := h.authService.GenerateToken(uint(newUser.ID), newUser.Email)
	if err != nil {
		response := helpers.JSONResult400{
			Message: "Register account failed",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.JSONResult200{
		Message: "Account has been registered",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    token,
	}
	c.JSON(http.StatusOK, response)
}

// UserLogin godoc
// @Summary      user login
// @Description  get token for login
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param		 login	body	user.LoginInput	true "Auth Login"
// @Success      200  {object} 	helpers.JSONResult200
// @Failure		 422  {object}  helpers.JSONResult422
// @Failure		 400  {object}  helpers.JSONResult400
// @Router       /api/login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.JSONResult422{
			Message: "Login failed",
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		response := helpers.JSONResult422{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Errors:  helpers.FormatValidationError(err),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(uint(loggedinUser.ID), loggedinUser.Email)
	if err != nil {
		response := helpers.JSONResult400{
			Message: "Login failed",
			Code:    http.StatusBadRequest,
			Status:  "error",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.JSONResult200{
		Message: "Successfuly logged",
		Code:    http.StatusOK,
		Status:  "success",
		Data:    token,
	}
	c.JSON(http.StatusOK, response)
}
