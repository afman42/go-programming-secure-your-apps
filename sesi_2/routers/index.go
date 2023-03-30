package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func LoadRouters(db *pgx.Conn) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	SetupUserRouters(r)
	SetupProductRouters(r)

	return r
}
