package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetAllData(url string) {
	resp, err := http.Get(url)

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
	var url = "https://jsonplaceholder.typicode.com/posts"
	GetAllData(url)
}
