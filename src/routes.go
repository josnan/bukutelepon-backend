package main

import "github.com/gin-gonic/gin"

type kontak struct {
	nama    string
	telepon string
}

// Routes of all Routes
func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"data": []gin.H{
				gin.H{
					"nama":    "Hadi Hidayat Hammurabi",
					"telepon": "082140320292",
				},
			},
		})
	})
}
