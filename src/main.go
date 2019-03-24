package main

import (
	"fmt"

	"./models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbconfig := config().dbconfig
	models.Init(fmt.Sprintf("postgresql://%s@%s:26257/%s?sslmode=disable", dbconfig.user, dbconfig.host, dbconfig.database))

	r := gin.Default()
	r.Use(corsMiddleware())
	Routes(r)

	r.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return cors.Default()
}
