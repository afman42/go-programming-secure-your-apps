package routers

import (
	"sesi_2_authentication_middleware/controllers"
	"sesi_2_authentication_middleware/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupProductRouters(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.Use(middlewares.Authentication())

		productAdminRouter := products.Group("/admin")
		productAdminRouter.GET("/", controllers.AllProducts)
		productAdminRouter.GET("/:productId", controllers.GetByIdProduct)
		productAdminRouter.POST("/", controllers.CreateProducts)
		productAdminRouter.PUT("/:productId", controllers.EditProduct)

		productUserRouter := products.Group("/user")
		productUserRouter.GET("/:productId", middlewares.ProductAuthentication(), controllers.GetByIdProduct)
		// productUserRouter.GET("/", middlewares.ProductAuthentication(), controllers.AllProducts)
		productUserRouter.POST("/", middlewares.ProductAuthentication(), controllers.CreateProducts)
	}
}
