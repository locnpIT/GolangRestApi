package main

import (
	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/db"
	"loc.com/hocgolang/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
