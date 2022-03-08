go-esa
===

[![codecov](https://codecov.io/gh/michimani/go-esa/branch/main/graph/badge.svg?token=JL9T8F4GBX)](https://codecov.io/gh/michimani/go-esa)  

Unofficial esa SDK for the Go programming language.

- [esa - 自律的なチームのための情報共有サービス](https://esa.io/)
- [dev/esa/api/v1 #noexpand - docs.esa.io](https://docs.esa.io/posts/102)

# Supported APIs

Progress of supporting APIs...

- **OAuth - 認証と認可**
  - `GET /oauth/token/info`
- **Teams - チーム**
  - `GET /v1/teams`
  - `GET /v1/teams/:team_name`
- **Stats - 統計情報**
  - `GET /v1/teams/:team_name/stats`

# Sample

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/go-esa/esaapi/oauth"
	"github.com/michimani/go-esa/esaapi/oauth/types"
	"github.com/michimani/go-esa/gesa"
)

func main() {
	token := "your-access-token"
	c, err := gesa.NewGesaClient(&gesa.NewGesaClientInput{
		AccessToken: token,
	})

	if err != nil {
		panic(err)
	}

	oauthTokenInfoGet(c)
}

func oauthTokenInfoGet(c *gesa.GesaClient) {
	r, err := oauth.OAuthTokenInfoGet(context.Background(), c, &types.OAuthTokenInfoGetParam{})
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

# Licence

[MIT](https://github.com/michimani/go-esa/blob/main/LICENCE)

# Author

[michimani210](https://twitter.com/michimani210)

