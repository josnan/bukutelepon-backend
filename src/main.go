package main

import (
	"./models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	models.Init("postgresql://root@128.199.191.240:26257/bukutelepon?sslmode=disable")

	r := gin.Default()
	Routes(r)

	r.Run()
}
