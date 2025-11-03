package model

import "time"

type Post struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Slug string `json:"slug"`
    Body string `json:"body"`
    Author string `json:"author"`
    CreatedAt time.Time `json:"created_at"`
}