package blogposts

import (
	"testing/fstest"
)

type Post struct {
	Title, Description, Body string
	Tags []string
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return nil
}
