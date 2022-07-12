package main

import "time"

type Recipe struct {
	Name        string    `json:"name"`
	Tags        []string  `json:"tags"`
	Ingradients []string  `json:"ingradients"`
	PublishedAt time.Time `json:"publishedAt"`
}
