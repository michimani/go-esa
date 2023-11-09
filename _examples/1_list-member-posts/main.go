package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/michimani/go-esa/v2/esaapi/models"
	"github.com/michimani/go-esa/v2/esaapi/post"
	"github.com/michimani/go-esa/v2/esaapi/post/types"
	"github.com/michimani/go-esa/v2/gesa"

	"github.com/olekukonko/tablewriter"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("The 1st (team_name), and 2nd (screen_name) parameters are is required.")
		os.Exit(1)
	}

	teamName := args[1]
	screenName := args[2]

	gesaClient, err := gesa.NewClient(&gesa.NewClientInput{
		AccessToken: os.Getenv("ESA_ACCESS_TOKEN"),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	posts, err := listPostsByMember(context.TODO(), gesaClient, teamName, screenName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Posts created by", screenName)
	showPostList(posts)
}

// listPostsByMember returns posts created by {screenName} in {teamName}.
// Returns the 10 most recently created.
func listPostsByMember(ctx context.Context, client *gesa.Client, teamName, screenName string) ([]models.Post, error) {
	in := &types.ListPostsInput{
		TeamName: teamName,
		Q:        fmt.Sprintf("user:%s", screenName),
		Sort:     types.ListPostsSortCreated,
		Order:    types.ListPostsOrderDesc,
		PerPage:  gesa.NewPageNumber(10),
	}
	out, err := post.ListPosts(ctx, client, in)
	if err != nil {
		return nil, err
	}

	return out.Posts, nil
}

const dateFmt = time.RFC3339

func showPostList(posts []models.Post) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Title", "URL", "Created At", "Updated At"})

	for i, p := range posts {
		table.Append([]string{
			strconv.Itoa(i + 1),
			p.Name,
			p.URL,
			p.CreatedAt.In(time.Local).Format(dateFmt),
			p.UpdatedAt.In(time.Local).Format(dateFmt),
		})
	}

	table.Render()
}
