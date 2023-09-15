package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetDataID3(url string) {
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
	url := "https://jsonplaceholder.typicode.com/posts/3"
	GetDataID3(url)
}
