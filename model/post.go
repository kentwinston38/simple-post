package model

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostReturn struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Post  `json:"data"`
}

type PostBatchReturn struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    []*Post `json:"data"`
}
