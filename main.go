package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbUser = os.Getenv("DB_USER")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")

func main() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
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

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Error loading DB")
	}

	return db
}
