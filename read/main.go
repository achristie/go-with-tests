package blogposts

import (
	"bufio"
	"io"
	"fmt"
	"bytes"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagSeperator				 = "Tags: "
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	var posts []Post

	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[len(titleSeperator):]
	description := readLine()[len(descriptionSeperator):]
	tags := strings.Split(readLine()[len(tagSeperator):], ", ")

	// ignore a line for the seperator
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")

	post := Post{Title: title, Description: description, Tags: tags, Body: body}
	return post, nil
}
