package controllers

import (
	"net/http"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/input"
	"sesi_2_authentication_middleware/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllUserByProducts(c *gin.Context) {
	models, err := models.AllByUserIDProducts(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err.Error(),
			"data":   models,
		})
		return
	}

	c.JSON(http.StatusOK, models)
}

func AllProducts(c *gin.Context) {
	models, err := models.AllProducts(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err.Error(),
			"data":   models,
		})
		return
	}

	c.JSON(http.StatusOK, models)
}

func GetByIdProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  "Invalid parameter",
		})
		return
	}
	model, err := models.GetByIdProduct(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not found",
			"error":  err.Error(),
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
			"status": "bad request",
			"error":  "Invalid parameter",
		})
		return
	}
	_, err = models.GetByIdProduct(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not found",
			"error":  err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"errors": helpers.FormatValidationError(err),
		})
		return
	}
	model, err := models.UpdateProductByID(c, inputProduct, uint(productID))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model)
}

func CreateProducts(c *gin.Context) {
	inputProduct := input.CreateOrUpdateProduct{}

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
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
			"status": "bad request",
			"error":  "Invalid parameter",
		})
		return
	}
	_, err = models.GetByIdProduct(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not found",
			"error":  err.Error(),
		})
		return
	}
	err = models.DeleteProductByID(c, uint(productID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sucessfull deleting data"})
}
