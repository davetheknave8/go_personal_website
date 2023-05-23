package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getProjects)

	router.Run("localhost:8080")
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sendProjects())
}

func postProjects(c *gin.Context) {
	var newProject project

	err := c.BindJSON(&newProject)
	if err != nil {
		return
	}

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}
