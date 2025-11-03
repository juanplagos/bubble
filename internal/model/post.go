package model

import "time"

type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Slug      string    `json:"slug"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    CreatedAt time.Time `json:"created_at"`
}