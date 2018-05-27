package main

import (
	"fmt"
	"os"
)

func main() {
	baseUrl := os.Getenv("API_ENDPOINT")
	token := os.Getenv("API_TOKEN")

	client, err := NewClient(baseUrl, token, &ClientConfig{
		TimeoutInSecond: 2,
	})
	if err != nil {
		os.Exit(1)
	}

	path := "/user/hirakiuc"
	user := ""

	pages, err := client.FetchPages(path, user)
	if err != nil {
		os.Exit(1)
	}

	for _, page := range pages {
		json, err := formatPage(page)
		if err != nil {
			continue
		}

		fmt.Println(json)
	}
}
