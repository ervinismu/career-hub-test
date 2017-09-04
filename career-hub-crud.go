package main

import (
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error


func main() {
 // NOTE: See we’re using = to assign the global var
 // instead of := which would assign it only in this function
 db, err = gorm.Open("sqlite3", "./gorm.db")
 if err != nil {
    fmt.Println(err)
 }
 defer db.Close()
 db.AutoMigrate(&Person{})
 r := gin.Default()

 r.Run(":8080")
}

func DeletePerson (c *gin.Context) {
	
	var person Person
	id := c.Params.ByName("id")

	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}

func UpdatePerson (c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; 
	err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)
}

func CreatePerson (c *gin.Context) {
	
	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(300, person)
}

func GetPerson(c *gin.Context) {
 id := c.Params.ByName("id")
 var person Person
 if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, person)
 }
}

func GetPeople(c *gin.Context) {
 var people []Person
 if err := db.Find(&people).Error; err != nil {
 c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, people)
 }
}
