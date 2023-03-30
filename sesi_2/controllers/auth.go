package controllers

import (
	"net/http"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/input"
	"sesi_2_authentication_middleware/models"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	inputRegister := input.RegisterUser{}

	if err := c.ShouldBindJSON(&inputRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": helpers.FormatValidationError(err),
		})
		return
	}

	user, err := models.RegisterUser(inputRegister, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func UserLogin(c *gin.Context) {
	inputUser := input.LoginUser{}
	password := ""

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": helpers.FormatValidationError(err),
		})
		return
	}

	password = inputUser.Password

	user, err := models.LoginUser(inputUser, c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	comparePass := helpers.CompareHash([]byte(user.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Password not same",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email, user.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
