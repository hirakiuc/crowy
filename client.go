package main

import (
	"fmt"
	"log"
	"time"

	crowi "github.com/crowi/go-crowi"
	"golang.org/x/net/context"
)

type Client struct {
	*crowi.Client
	Config *ClientConfig
}

type ClientConfig struct {
	TimeoutInSecond time.Duration
}

func defaultConfig() *ClientConfig {
	return &ClientConfig{
		TimeoutInSecond: 2,
	}
}

func NewClient(url string, token string, config *ClientConfig) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("You should pass a toke string.")
	}

	crowiConfig := crowi.Config{
		URL:   url,
		Token: token,
	}
	client, err := crowi.NewClient(crowiConfig)
	if err != nil {
		log.Printf("[ERROR] failed to get client. %s", err.Error())
		return nil, err
	}

	if config != nil {
		return &Client{client, config}, nil
	} else {
		return &Client{client, defaultConfig()}, nil
	}
}

func (client *Client) FetchPages(path string, user string) ([]crowi.PageInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.Config.TimeoutInSecond*time.Second)
	defer cancel()

	opts := &crowi.PagesListOptions{
		ListOptions: crowi.ListOptions{
			Pagenation: true,
		},
	}

	res, err := client.Pages.List(ctx, path, user, opts)
	if err != nil {
		log.Printf("[ERROR] failed to fetch pages: %s", err.Error())
		return []crowi.PageInfo{}, err
	}

	if !res.OK {
		log.Printf("[ERROR] Error returned: %s", res.Error)
		return []crowi.PageInfo{}, err
	}

	return res.Pages, nil
}
