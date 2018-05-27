package main

import (
	"encoding/json"
	"fmt"

	crowi "github.com/crowi/go-crowi"
)

type WorkflowItem struct {
	Valid        bool
	UUID         string
	Title        string
	SubTitle     string
	Arg          string
	AutoComplete string
	QuickLookURL string
	Text         string
}

func formatPage(pageInfo crowi.PageInfo) (string, error) {
	item := WorkflowItem{
		Valid:        true,
		UUID:         pageInfo.ID,
		Title:        pageInfo.Path,
		SubTitle:     pageInfo.Status,
		Arg:          "",
		AutoComplete: "auto-completed",
		QuickLookURL: "",
		Text:         pageInfo.ID,
	}

	bin, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(bin), nil
}
