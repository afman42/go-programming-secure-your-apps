package controllers

import (
	"net/http"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/input"
	"sesi_2_authentication_middleware/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllProducts(c *gin.Context) {
	models, err := models.AllProducts(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models)
}

func GetByIdProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	model, err := models.GetByIdProduct(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model)
}

func EditProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	inputProduct := input.CreateOrUpdateProduct{}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	GetByIdProduct(c)

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": helpers.FormatValidationError(err),
		})
		return
	}
	model, err := models.UpdateProductByID(c, inputProduct, uint(productID))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model)
}

func CreateProducts(c *gin.Context) {
	inputProduct := input.CreateOrUpdateProduct{}

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": helpers.FormatValidationError(err),
		})
		return
	}

	model, err := models.CreateProduct(c, inputProduct)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, model)
}

func DeleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	GetByIdProduct(c)
	err = models.DeleteProductByID(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sucessfull deleting data"})
}
