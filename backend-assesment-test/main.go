package main

import (
	"os"

	"github.com/ervinismu/backend-assesment-test/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.GET("/notes")
	r.POST("/notes", controllers.CreateNotes)
	r.PUT("/notes/:id", controllers.UpdateNotes)
	r.DELETE("/notes/:id", controllers.DeleteNotes)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/logout", controllers.Logout)
	r.Run(":" + port)
}
