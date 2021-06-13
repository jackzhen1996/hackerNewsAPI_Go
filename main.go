package main

import (
	"fmt"
	"net/http"
	"log"
	// "hackerNewsAPI/routes"
	"hackerNewsAPI/controller"
)

func main() {
	fmt.Printf("serving at localhost:8080")
	// http.HandleFunc("/", routes.Router)
	go http.HandleFunc("/getAllFeeds", controller.GetAllFeeds)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
