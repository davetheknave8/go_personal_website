package main

// project represents a personal project developed or contributed by me
type project struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
}
