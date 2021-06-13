package routes

import (
	"fmt"
	"net/http"
	"hackerNewsAPI/controller"
)

// Takes endpoint and callback as argument a nd runs handleFunc
func route(endpoint string, callback func(w http.ResponseWriter, r*http.Request)) {
	http.HandleFunc(endpoint, callback)
}

// Router that redirects traffic to different controller functions
func Router(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in router")

	// get all feeds
	route("/getAllFeeds", controller.GetAllFeeds)
}

