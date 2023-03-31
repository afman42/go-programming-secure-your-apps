package middlewares

import (
	"net/http"
	"sesi_2_authentication_middleware/enums"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unautenticated",
				"message": err.Error(),
			})
			return
		}
		verRoleUser := verifyToken.(*helpers.RoleUserClaims).Role
		if !strings.Contains(c.Request.URL.Path, string(verRoleUser)) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "Your Role is not allowed",
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

func ProductAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*pgx.Conn)
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "bad request",
				"messaage": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(*helpers.RoleUserClaims)
		userID := uint(userData.ID)

		product, err := models.GetOneProductByUserId(db, uint(productId))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data does'nt exist",
			})
			return
		}

		if userData.Role == enums.User {
			if product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}
		c.Next()
	}
}
