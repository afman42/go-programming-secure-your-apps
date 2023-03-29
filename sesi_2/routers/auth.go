package routers

import (
	"sesi_2_authentication_middleware/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRouters(r *gin.Engine) {
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
}
