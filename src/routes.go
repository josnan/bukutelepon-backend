package main

import (
	"log"

	"./models"
	"github.com/gin-gonic/gin"
)

type kontak struct {
	nama    string
	telepon string
}

// Routes of all Routes
func Routes(r *gin.Engine) {
	r.GET("/", index)
}

func index(c *gin.Context) {
	contacts, err := models.ContactModelAll()

	if err != nil {
		log.Fatal("Model Error Get All Contacts: ", err)
	}

	c.JSON(200, contacts)
}
