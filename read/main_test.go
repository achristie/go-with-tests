package blogposts_test

import (
	blogposts "github.com/achristie/go-with-tests/read"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md": {Data: []byte("hi")},
		"hello_world2.md":{Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
