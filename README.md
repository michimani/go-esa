go-esa
===

[![Go Reference](https://pkg.go.dev/badge/github.com/michimani/go-esa.svg)](https://pkg.go.dev/github.com/michimani/go-esa)
[![codecov](https://codecov.io/gh/michimani/go-esa/branch/main/graph/badge.svg?token=JL9T8F4GBX)](https://codecov.io/gh/michimani/go-esa)  

Unofficial esa SDK for the Go programming language.

- [esa - 自律的なチームのための情報共有サービス](https://esa.io/)
- [dev/esa/api/v1 #noexpand - docs.esa.io](https://docs.esa.io/posts/102)

# Supported APIs

Progress of supporting APIs...

- **OAuth**
  - `GET /oauth/token/info`
- **Team**
  - `GET /v1/teams`
  - `GET /v1/teams/:team_name`
- **Stats**
  - `GET /v1/teams/:team_name/stats`
- **Member**
  - `GET /v1/teams/:team_name/members`
  - `DELETE /v1/teams/:team_name/members/:screen_name`
- **Post**
  - `GET /v1/teams/docs/posts`
  - `GET /v1/teams/:team_name/posts/:post_number`
  - `POST /v1/teams/:team_name/posts`
  - `PATCH /v1/teams/:team_name/posts/:post_number`
  - `DELETE /v1/teams/:team_name/posts/:post_number`
- **Comment**
  - `GET /v1/teams/:team_name/posts/:post_number/comments`
  - `GET /v1/teams/:team_name/comments/:comment_id`
  - `POST /v1/teams/:team_name/posts/:post_number/comments`
  - `PATCH /v1/teams/:team_name/comments/:comment_id`
  - `DELETE /v1/teams/:team_name/comments/:comment_id`
  - `GET /v1/teams/:team_name/comments`
- **Star**
  - `GET /v1/teams/:team_name/posts/:post_number/stargazers`
  - `POST /v1/teams/:team_name/posts/:post_number/star`
  - `DELETE /v1/teams/:team_name/posts/:post_number/star`
  - `GET /v1/teams/:team_name/comments/:comment_id/stargazers`
  - `POST /v1/teams/:team_name/comments/:comment_id/star`
  - `DELETE /v1/teams/:team_name/comments/:comment_id/star`
- **Watch**
  - `GET /v1/teams/:team_name/posts/:post_number/watchers`
  - `POST /v1/teams/:team_name/posts/:post_number/watch`
  - `DELETE /v1/teams/:team_name/posts/:post_number/watch`
- **Tag**
  - `GET v1/teams/:team_name/tags`
- **Category**
  - `POST v1/teams/:team_name/categories/batch_move`
- **Emoji**
  - `GET /v1/teams/:team_name/emojis`
  - `POST /v1/teams/:team_name/emojis`
  - `DELETE /v1/teams/:team_name/emojis/:code`
- **User**
  - `GET /v1/user`
- **Invitation**
  - `GET /v1/teams/:team_name/invitation`

# Sample

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/go-esa/esaapi/oauthtoken"
	"github.com/michimani/go-esa/esaapi/oauthtoken/types"
	"github.com/michimani/go-esa/gesa"
)

func main() {
	token := "your-access-token"
	c, err := gesa.NewClient(&gesa.NewClientInput{
		AccessToken: token,
	})

	if err != nil {
		panic(err)
	}

	getOAuthTokenInfo(c)
}

func getOAuthTokenInfo(c *gesa.Client) {
	r, err := oauthtoken.GetOAuthTokenInfo(context.Background(), c, &types.GetOAuthTokenInfoInput{})
	if err != nil {
		fmt.Println(err)

		ge := err.(*gesa.GesaError)
		if ge.OnAPI {
			fmt.Println(ge.EsaAPIError.StatusCode)
			fmt.Println(ge.EsaAPIError.Status)
			fmt.Println(ge.EsaAPIError.Error)
			fmt.Println(ge.EsaAPIError.Message)

			if ge.EsaAPIError.RateLimitInfo != nil {
				fmt.Println(ge.EsaAPIError.RateLimitInfo.Limit)
				fmt.Println(ge.EsaAPIError.RateLimitInfo.Remaining)
				fmt.Println(ge.EsaAPIError.RateLimitInfo.Reset)
			}
		}

		return
	}

	fmt.Printf("Response: %+v \n", r)

	fmt.Println(gesa.IntValue(r.ResourceOwnerID))
	fmt.Println(r.Scope)
	fmt.Println(r.Application.UID)
	fmt.Println(r.CreatedAt.Time())
	fmt.Println(r.User.ID)
	fmt.Println(r.RateLimitInfo)
}
```

# License

[MIT](https://github.com/michimani/go-esa/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)

