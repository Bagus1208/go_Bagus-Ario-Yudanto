package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type Cloth []struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	var wait sync.WaitGroup
	var cloth Cloth

	wait.Add(1)
	go func() {
		defer wait.Done()

		response, err := http.Get("https://fakestoreapi.com/products")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		defer response.Body.Close()
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(data, &cloth)
	}()

	wait.Wait()

	for i := 0; i < len(cloth); i++ {
		fmt.Println("===")
		fmt.Println("Title : ", cloth[i].Title)
		fmt.Println("Price : ", cloth[i].Price)
		fmt.Println("Category : ", cloth[i].Category)
		fmt.Println("===")
	}
}
