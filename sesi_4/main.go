package main

import (
	"sesi_4_final_project/auth"
	"sesi_4_final_project/database"
	"sesi_4_final_project/handler"
	"sesi_4_final_project/middlewares"
	"sesi_4_final_project/socialmedia"
	"sesi_4_final_project/user"

	"sesi_4_final_project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	docs.SwaggerInfo.Title = "Swagger Example MyGram API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := database.LoadDB()
	userRepository := user.NewRepository(db)
	socialMediaRepository := socialmedia.NewRepository(db)

	userService := user.NewService(userRepository)
	socialMediaService := socialmedia.NewService(socialMediaRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	api.POST("/login", userHandler.Login)
	api.POST("/register", userHandler.RegisterUser)

	socialMediaApi := api.Group("/social_media")
	socialMediaApi.GET("/", middlewares.AuthMiddleware(authService, userService), socialMediaHandler.GetAllSocialMedia)
	r.Run(":8080")
}
