package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/projects/", getProjects).Methods("GET")
	router.HandleFunc("/projects", postProjects).Methods("POST")
	router.HandleFunc("/projects/:id", getProjectById).Methods("GET")
}

func getProjects(w http.ResponseWriter, r *http.Request) {
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

func getProjectById(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	fmt.Println("Getting project")

	row, err := db.Query("SELECT * FROM PROJECTS WHERE id = %d", r.FormValue("projectid"))
	if err != nil {
		log.Fatalf("Unable to retrieve project")
	}

	var id int
	var title string
	var description string
	var links []string

	err = row.Scan(&id, &title, &description, &links)

	response := JsonResponse{Type: "success", Data: project{ID: id}}
}

func postProjects(w http.ResponseWriter, r *http.Request) {
	projectID := r.FormValue("projectid")
	projectName := r.FormValue("projectname")

	var response = JsonResponse{}

	if projectID == "" || projectName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing a ProjectID or a ProjectName parameter."}
	} else {
		db := setupDB()

		fmt.Println("Inserting movie into DB")

		fmt.Println("Inserting new Project with ID: " + projectID + " and name: " + projectName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO projects(projectID, projectName) VALUES($1, $2) returning id;", projectID, projectName).Scan(&lastInsertID)

		// check errors
		if err != nil {
			log.Fatalf("Error inserting project into DB")
		}

		response = JsonResponse{Type: "success", Message: "The project has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
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
