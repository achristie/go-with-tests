package blogposts_test

import (
	"read"
	"testing"
	"testing/fstest"
	"reflect"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md": {Data: []byte("Title: Post 1")},
		"hello_world2.md":{Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})

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
