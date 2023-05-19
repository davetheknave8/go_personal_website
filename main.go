package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(sendProjects())
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sendProjects())
}
