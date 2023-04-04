package main

import (
	"sesi_4_final_project/auth"
	"sesi_4_final_project/comments"
	"sesi_4_final_project/database"
	"sesi_4_final_project/handler"
	"sesi_4_final_project/middlewares"
	"sesi_4_final_project/photo"
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
	commentsRepository := comments.NewRepository(db)
	photoRepository := photo.NewRepository(db)

	userService := user.NewService(userRepository)
	socialMediaService := socialmedia.NewService(socialMediaRepository)
	commentsService := comments.NewService(commentsRepository)
	photoService := photo.NewService(photoRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)
	commentsHandler := handler.NewCommentsHandler(commentsService)
	photoHandler := handler.NewPhotoHandler(photoService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	api.POST("/login", userHandler.Login)
	api.POST("/register", userHandler.RegisterUser)

	socialMediaApi := api.Group("/social_media")
	socialMediaApi.GET("/", middlewares.AuthMiddleware(authService, userService), socialMediaHandler.GetAllSocialMedia)
	socialMediaApi.GET("/:socialMediaID", middlewares.AuthMiddleware(authService, userService), socialMediaHandler.GetOneSocialMedia)
	socialMediaApi.POST("/", middlewares.AuthMiddleware(authService, userService), socialMediaHandler.CreateSocialMedia)
	socialMediaApi.PUT("/:socialMediaID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationSocialMediaByUserID(socialMediaService), socialMediaHandler.UpdateSocialMedia)
	socialMediaApi.DELETE("/:socialMediaID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationSocialMediaByUserID(socialMediaService), socialMediaHandler.DeleteSocialMedia)

	commentApi := api.Group("/comment/:photoID")
	commentApi.GET("/", middlewares.AuthMiddleware(authService, userService), commentsHandler.GetAllComment)
	commentApi.GET("/:commentID", middlewares.AuthMiddleware(authService, userService), commentsHandler.GetOneComment)
	commentApi.POST("/", middlewares.AuthMiddleware(authService, userService), commentsHandler.CreateComment)
	commentApi.PUT("/:commentID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationCommentByUserID(commentsService), commentsHandler.UpdateComment)
	commentApi.DELETE("/:commentID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationCommentByUserID(commentsService), commentsHandler.DeleteComment)

	photoApi := api.Group("/photo")
	photoApi.GET("/", middlewares.AuthMiddleware(authService, userService), photoHandler.GetAllPhoto)
	photoApi.GET("/:photoID", middlewares.AuthMiddleware(authService, userService), photoHandler.GetOnePhoto)
	photoApi.POST("/", middlewares.AuthMiddleware(authService, userService), photoHandler.CreatePhoto)
	photoApi.PUT("/:photoID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationPhotoByUserID(photoService), photoHandler.UpdatePhoto)
	photoApi.DELETE("/:photoID", middlewares.AuthMiddleware(authService, userService), middlewares.AuthorizationPhotoByUserID(photoService), photoHandler.DeletePhoto)

	r.Run(":8080")
}
