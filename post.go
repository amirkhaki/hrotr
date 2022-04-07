package main

type post struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Hash string `json:"hash"`
	Description string `json:"description"`
	FDescription string `json:"full_description"`
	Image string `json:"image"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
}
