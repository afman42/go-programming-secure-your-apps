package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRouters(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	SetupUserRouters(r)
	SetupProductRouters(r)

	return r
}
