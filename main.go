package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, rt *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func indexArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: return all articles")
	json.NewEncoder(w).Encode(Articles)
}

func showArticle(w http.ResponseWriter, r *http.Request) {
	// := shorthand variable declaration
	// available inside function scope; otherwise use 'var'
	params := mux.Vars(r)
	// Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
	for _, article := range Articles {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
   // because readAll returns 2 vals
   reqBody, _ := ioutil.ReadAll(r.Body)
   fmt.Println("Endpoint hit: creating article with params:", string(reqBody))

   // initialize article
   // unmarshal JSON into a struct.
   // match incoming JSON fields to the keys used by Marshal, prefer exact match but case insensitive
   var article Article
   json.Unmarshal(reqBody, &article)

   Articles = append(Articles, article)
   fmt.Fprint(w, "%+v", string(reqBody))

   json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter,  r *http.Request) {
	// path params from mux
	params := mux.Vars(r)

	id := params["id"]

	for idx, article := range Articles {
	    if article.ID == id {
			// curious about this syntax, spread?
            // what abt the colon after idx + 1?
			Articles = append(Articles[:idx], Articles[idx+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Println("Endpoint hit: updating article with params:", string(reqBody))
	
	params := mux.Vars(r)
    id := params["id"]

	var articleParams Article
	json.Unmarshal(reqBody, &articleParams)

	for idx, article := range Articles {
        if article.ID == id {
		  fmt.Println("This article right here: ", article, idx)
		  // replace current article with article params
		  Articles[idx] =  articleParams
		}
	}
    fmt.Println("updated articles list:", Articles)
}

func handleRequests() {
	// initialize muxrouter
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/articles", createArticle).Methods("POST")
	muxRouter.HandleFunc("/articles", indexArticles)
	muxRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	muxRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	muxRouter.HandleFunc("/articles/{id}", showArticle)
	// pass router instance to server
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{ID: "1", Title: "Article 1", Desc: "first description", Content: "Blah Blah"},
		Article{ID: "2", Title: "Article 2", Desc: "second description", Content: "Blah Blah"},
	}
	handleRequests()
}

// Article is...
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

// Articles is ...
var Articles []Article
