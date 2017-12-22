package controllers

import (
	"fmt"

	"github.com/ervinismu/backend-assesment-test/models"
	"github.com/gin-gonic/gin"
)

//CreateNotes used for create a new notes
func CreateNotes(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var notes models.Notes
	c.Bind(&notes)
	var res models.Response
	c.Bind(&notes)
	if err := db.Create(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		res.Code = "200"
		res.Message = "Post Success!"
		c.JSON(200, res)
	}
}

//DeleteNotes used for delete notes that exist before
func DeleteNotes(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var notes models.Notes
	var res models.Response
	db.Where("id = ?", id).First(&notes)
	if err := db.Delete(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(404)
	} else {
		res.Code = "200"
		res.Message = "Notes " + id + " has been deleted."
		c.JSON(200, res)
	}
}

//UpdateNotes used for edit post
func UpdateNotes(c *gin.Context) {
	var res models.Response
	var notes models.Notes
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&notes)
	db.Save(&notes)
	res.Code = "200"
	res.Message = "Update post success!"
	c.JSON(200, res)
}
