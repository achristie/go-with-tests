package main

import (
	blogposts "github.com/achristie/go-with-tests/read"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range posts {
		log.Printf("Title: %v, \nDescription: %v, \nTags: %v, \nBody: %v",
				s.Title, s.Description, s.Tags, s.Body)
	}

}
