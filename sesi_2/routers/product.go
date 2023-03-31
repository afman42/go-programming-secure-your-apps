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
		productAdminRouter.GET("/:productID", controllers.GetByIdProduct)
		productAdminRouter.POST("/", controllers.CreateProducts)
		productAdminRouter.PUT("/:productID", controllers.EditProduct)
		productAdminRouter.DELETE("/:productID", controllers.DeleteProduct)

		productUserRouter := products.Group("/user")
		productUserRouter.GET("/:productID", middlewares.ProductAuthentication(), controllers.GetByIdProduct)
		productUserRouter.GET("/", controllers.AllUserByProducts)
		productUserRouter.POST("/", controllers.CreateProducts)
	}
}
