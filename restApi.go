package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Title 1",
			Desc:    "Desc 1",
			Content: "Content 1",
		},
		Article{
			Title:   "Title 2",
			Desc:    "Desc 2",
			Content: "Content 2",
		},
		Article{
			Title:   "Title 3",
			Desc:    "Desc 3",
			Content: "Content 3",
		},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage endpoint Hit")
	json.NewEncoder(w).Encode("Hello World")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
