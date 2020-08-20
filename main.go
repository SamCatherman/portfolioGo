package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, rt *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return all articles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{Title: "Article 1", Desc: "first description", Content: "Blah Blah"},
		Article{Title: "Article 2", Desc: "second description", Content: "Blah Blah"},
	}
	handleRequests()
}

// Article is...
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

// Articles is ...
var Articles []Article
