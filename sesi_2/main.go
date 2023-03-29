package main

import (
	"sesi_2_authentication_middleware/database"
	"sesi_2_authentication_middleware/routers"
)

func main() {
	db := database.LoadDB()
	routers.LoadRouters(db).Run(":8080")
}
