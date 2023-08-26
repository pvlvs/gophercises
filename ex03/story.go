package main

import (
	"encoding/json"
	"html"
	"os"
)

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type storyArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

func storyMap() map[string]storyArc {
	f, err := os.ReadFile("gopher.json")
    check(err)

	sa := make(map[string]storyArc)
	json.Unmarshal(f, &sa)

	for _, s := range sa {
		s.Title = html.EscapeString(s.Title)

		for j := range s.Story {
			s.Story[j] = html.EscapeString(s.Story[j])
		}
		for j := range s.Options {
			s.Options[j].Text = html.EscapeString(s.Options[j].Text)
		}
	}

	return sa
}
