package main

import (
	"firstGolangProject/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/articles", routes.AllArticles)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
