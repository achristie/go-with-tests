package main

import (
	blogposts "github.com/achristie/go-with-tests/read"
	"log"
	"os"
)

func main() {
	log.Println(os.DirFS("posts"))
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range posts {
		log.Printf("Title: %v, Description: %v, Tags: %v, Body: %v",
				s.Title, s.Description, s.Tags, s.Body)
	}

}
