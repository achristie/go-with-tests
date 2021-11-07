package blogposts_test

import (
	blogposts "github.com/achristie/go-with-tests/read"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: tag3, tag4, tag5
---
Test
That
Thing`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertPost(t, posts[0], blogposts.Post{
									Title: "Post 1", 
									Description: "Description 1",
									Tags: []string{"tag1", "tag2"},
									Body: "Hello\nWorld"})

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
