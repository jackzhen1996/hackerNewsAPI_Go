package controller

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Feed struct {
	By string `json:"by"`
	Descendants int `json:"descendants"`
	Score int `json:"score"`
	Time int `json:"time"`
	Title string `json:"title"`
	Kids []int `json:"kids"`
	Url string `json:"url"`
}

type AllFeed struct {
	Feeds []Feed
}


// get single feed from id
func getFeed(w http.ResponseWriter, r *http.Request,id int) Feed {
	res,err := http.Get("https://hacker-news.firebaseio.com/v0/item/"+ strconv.Itoa(id) + ".json?" + "print=pretty")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Cannot get a feed")
	}

	bytes, err2 := ioutil.ReadAll(res.Body)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		log.Print("Cannot parse individual feed in byte form")
	}

	var feed Feed
	json.Unmarshal(bytes,&feed)
	return feed
}

// get all feed ids
func getAllFeedIds(w http.ResponseWriter, r *http.Request) []int  {
	var BaseURL = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

	fmt.Println("getting ids of all feeds")
	res, err := http.Get(BaseURL)

	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Cannot get feed from API")
	}

	bytes, err2 := ioutil.ReadAll(res.Body)

	if (err2 != nil) {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		log.Print("Cannot parse ids in byte form")
	}

	var idList []int
	json.Unmarshal(bytes, &idList)

	// return the list of ids
	return idList[0:21]
}

// get all feeds from each id
func GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all feeds")

	// get all the ids
	idList := getAllFeedIds(w,r)

	var feedsArray []Feed
	// for each id, make a request to retrieve the data
	for _, id := range idList {
		feedsArray = append(feedsArray, getFeed(w,r,id))
	}

	// send back to client
	json.NewEncoder(w).Encode(feedsArray)
}


