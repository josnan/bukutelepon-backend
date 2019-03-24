package main

import (
	"fmt"

	"./models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbconfig := config().dbconfig
	models.Init(fmt.Sprintf("postgresql://%s@%s:26257/%s?sslmode=disable", dbconfig.user, dbconfig.host, dbconfig.database))

	r := gin.Default()
	Routes(r)

	r.Run()
}
