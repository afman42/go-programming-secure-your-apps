package controllers

import (
	"net/http"
	"sesi_2_authentication_middleware/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllProducts(c *gin.Context) {

}

func GetByIdProduct(c *gin.Context) {

}

func EditProduct(c *gin.Context) {

}

func CreateProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userData := c.MustGet("userData").(jwt.MapClaims)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	c.ShouldBind(&Product)

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, Product)
}

func DeleteProduct(c *gin.Context) {

}
