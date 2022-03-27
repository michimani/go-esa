package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/michimani/go-esa/esaapi/post"
	"github.com/michimani/go-esa/esaapi/post/types"
	"github.com/michimani/go-esa/gesa"
)

const endCode = "EOF"

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("The 1st (team name), and 2nd (post title) parameters are is required.")
		os.Exit(1)
	}

	teamName := args[1]
	postTitle := args[2]

	gesaClient, err := gesa.NewClient(&gesa.NewClientInput{
		AccessToken: os.Getenv("ESA_ACCESS_TOKEN"),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("// Please enter the content. Enter \"EOF\" to confirm your post.")

	var sb strings.Builder
	s := bufio.NewScanner(os.Stdin)
	for {
		s.Scan()
		line := s.Text()
		if line == endCode {
			break
		}
		sb.WriteString(line + "\n")
	}

	url, err := createNewPost(context.TODO(), gesaClient, teamName, postTitle, sb.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("New post \"%s\" is successfully created. URL is %s\n", postTitle, url)
}

func createNewPost(ctx context.Context, client *gesa.Client, teamName, title, body string) (string, error) {
	in := &types.CreatePostInput{
		TeamName: teamName,
		Name:     title,
		BodyMD:   gesa.String(body),
	}

	out, err := post.CreatePost(ctx, client, in)
	if err != nil {
		return "", err
	}

	return out.URL, nil
}
