package main

import (
	"context"
	"sesi_2_authentication_middleware/database"
	"sesi_2_authentication_middleware/routers"
)

func main() {
	db := database.LoadDB()
	defer db.Close(context.Background())

	routers.LoadRouters(db).Run(":8080")
}
