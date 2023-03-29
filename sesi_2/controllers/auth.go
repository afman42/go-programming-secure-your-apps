package controllers

import (
	"net/http"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	User := models.User{}
	c.ShouldBindJSON(&User)

	// fmt.Println(User.Password)
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": helpers.FormatResponseValidation(err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    User.ID,
		"email": User.Email,
	})
}

func UserLogin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	User := models.User{}
	LoginUser := models.LoginUser{}
	password := ""

	c.ShouldBindJSON(&LoginUser)
	_, err := govalidator.ValidateStruct(LoginUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Unprocessableentity",
			"message": helpers.FormatResponseValidation(err.Error()),
		})
		return
	}

	password = LoginUser.Password

	err = db.Debug().First(&User, "email = ?", LoginUser.Email).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	comparePass := helpers.CompareHash([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalida email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
