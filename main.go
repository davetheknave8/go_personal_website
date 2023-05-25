package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/projects", getProjects)
	router.GET("/projects/:id", getProjectById)
	router.POST("/projects", postProjects)

	router.Run("localhost:8080")
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sendProjects())
}

func getProjectById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range sendProjects() {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
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
