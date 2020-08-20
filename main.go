package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, rt *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return all articles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	// initialize muxrouter
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/articles", returnAllArticles)
	// pass router instance to server
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
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
