package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (post *Post) PostData(url string) {
	jsonReq, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	var post = Post{
		UserID: 7,
		Title:  "Bagus Ario Yudanto",
		Body:   "Test post data ke",
	}

	post.PostData(url)
}

// package main

// import (
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// type Post struct {
// 	UserID string `json:"userId"`
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }

// func main() {
// 	params := url.Values{}
// 	params.Add("userId", "7")
// 	params.Add("title", "Bagus Ario Yudanto")
// 	params.Add("body", "Test post data ke API")

// 	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", params)
// 	if err != nil {
// 		log.Printf("Request Failed: %s", err)
// 		return
// 	}

// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.BodMarshal(v)

// 	// Log the request body
// 	bodyString := string(body)
// 	log.Print(bodyString)

// 	// Unmarshal result
// 	post := Post{}
// 	err = json.Unmarshal(body, &post)
// 	if err != nil {
// 		log.Printf("Reading body failed: %s", err)
// 		return
// 	}

// 	log.Printf("Post added with ID %d", post.ID)
// }
