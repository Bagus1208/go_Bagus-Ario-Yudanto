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

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(bodyBytes))
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/7"

	var post = Post{
		UserID: 7,
		Title:  "Bagus Ario Yudanto",
		Body:   "Test post data ke",
	}

	post.PostData(url)
}
