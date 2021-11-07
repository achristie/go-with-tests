package main

import (
	"log"
	"read"
	"os"
)

func main() {
	posts, err := read.NewPostsFromFS("/posts")

	if err != nil {
		log.Fatal(err)
	}

	for i, s := range posts {
		log.Printf("Title: %v, Description: %v, Tags: %v, Body: %v",
				s.Title, s.Description, s.Tags, s.Body)
	}

}
