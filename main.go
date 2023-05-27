package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbUser = os.Getenv("DB_USER")
var dbPassword = os.Getenv("DB_PASSWORD")
var dbName = os.Getenv("DB_NAME")

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []project `json:"data"`
	Message string    `json:"string"`
}

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

func getProjects(c *gin.Context, w http.ResponseWriter) {
	db := setupDB()

	fmt.Println("Getting projects...")

	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		log.Fatalf("Error loading projects from DB")
	}

	var projects []project

	for rows.Next() {
		var id int
		var title string
		var description string
		var links []string

		err = rows.Scan(&id, &title, &description, &links)

		if err != nil {
			fmt.Println("Error scanning rows in DB response for projects")
		}

		projects = append(projects, project{ID: id, Title: title, Description: description, Links: links})
	}

	var response = JsonResponse{Type: "success", Data: projects}

	json.NewEncoder(w).Encode(response)
}

func getProjectById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range sendProjects() {
		if strconv.Itoa(a.ID) == id {
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
