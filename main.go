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
