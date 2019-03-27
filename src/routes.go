package main

import (
	"log"

	"./models"
	"github.com/gin-gonic/gin"
)

// Response structure
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Routes of all Routes
func Routes(r *gin.Engine) {
	r.GET("/", index)
	r.POST("/", insertContact)
	r.DELETE("/:id", deleteContact)
	r.PUT("/:id", updateContact)
}

func index(c *gin.Context) {
	contacts, err := models.ContactModelAll()

	if err != nil {
		log.Fatal("Model Error Get All Contacts: ", err)
		response := &Response{
			Status:  "FAIL",
			Message: "Error while getting contacts.",
		}
		c.JSON(200, response)
	} else {
		c.JSON(200, contacts)
	}
}

func insertContact(c *gin.Context) {
	var newcontact *models.Contact
	var response *Response
	err := c.ShouldBindJSON(&newcontact)

	if err != nil {
		log.Fatal("Create Contact Error: ", err)
		response = &Response{
			Status:  "FAIL",
			Message: "Error while inserting contact.",
		}
	}

	success, err := models.ContactModelCreate(newcontact)

	if err != nil {
		log.Fatal("Create Contact Error: ", err)
		response = &Response{
			Status:  "FAIL",
			Message: "Error while inserting contact.",
		}
	}

	if success {

	}

	response = &Response{
		Status:  "OK",
		Message: "New contact has been saved.",
	}

	c.JSON(200, response)
}

func deleteContact(c *gin.Context) {
	var response *Response
	id := c.Param("id")

	_, err := models.ContactModelDelete(id)

	if err != nil {
		log.Fatal("Delete Contact Error: ", err)
		response = &Response{
			Status:  "FAIL",
			Message: "Error while deleting contact.",
		}
	}

	response = &Response{
		Status:  "OK",
		Message: "Contact has been deleted.",
	}

	c.JSON(200, response)
}

// func updateContact(c *gin.Context) {
// 	var response *Response
// 	var newcontact *models.Contact
// 	id := c.Param("id")

// 	contact, err := models.ContactModelByID(id)

// 	if err != nil || contact.ID == 0 {
// 		response = &Response{
// 			Status:  "FAIL",
// 			Message: "Contact not found.",
// 		}
// 		c.JSON(404, response)
// 	} else {

// 		_, err := models.ContactModelUpdate(id, )

// 		c.JSON(200, contact)
// 	}

// }
