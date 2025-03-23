package main

import (
	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/db"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.Run(":8080") //localhost:8080

}
