package main

// project represents a personal project developed or contributed by me
type project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
}

var projects = []project{
	{ID: "1", Title: "Sample Project", Description: "This is a sample project.", Links: []string{"www.test.com, www.sample.org"}},
}

func sendProjects() []project {
	return projects
}
