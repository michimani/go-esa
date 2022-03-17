// Package goesa is the unofficial esa API v1 SDK for the Go programming language.
//
// Getting started: To get information of OAuth access token.
//
//	go get github.com/michimani/go-esa/gesa
// 	go get github.com/michimani/go-esa/esaapi/oauth
//	go get github.com/michimani/go-esa/esaapi/oauth/types
//
//
//     package main
//
//     import (
//         "context"
//         "fmt"
//         "os"
//
//         "github.com/michimani/go-esa/esaapi/oauth"
//         "github.com/michimani/go-esa/esaapi/oauth/types"
//         "github.com/michimani/go-esa/gesa"
//     )
//
//     func main() {
//         token := "your-access-token"
//         c, err := gesa.NewGesaClient(&gesa.NewGesaClientInput{
//             AccessToken: token,
//         })
//
//         if err != nil {
//             panic(err)
//         }
//
//         oauthTokenInfoGet(c)
//     }
//
//     func oauthTokenInfoGet(c *gesa.GesaClient) {
//         r, err := oauth.OAuthTokenInfoGet(context.Background(), c, &types.OAuthTokenInfoGetParam{})
//         if err != nil {
//             fmt.Println(err)
//
//             ge := err.(*gesa.GesaError)
//             if ge.OnAPI {
//                 fmt.Println(ge.EsaAPIError.StatusCode)
//                 fmt.Println(ge.EsaAPIError.Status)
//                 fmt.Println(ge.EsaAPIError.Error)
//                 fmt.Println(ge.EsaAPIError.Message)
//
//                 if ge.EsaAPIError.RateLimitInfo != nil {
//                     fmt.Println(ge.EsaAPIError.RateLimitInfo.Limit)
//                     fmt.Println(ge.EsaAPIError.RateLimitInfo.Remaining)
//                     fmt.Println(ge.EsaAPIError.RateLimitInfo.Reset)
//                 }
//             }
//
//             return
//         }
//
//         fmt.Printf("Response: %+v \n", r)
//
//         fmt.Println(gesa.IntValue(r.ResourceOwnerID))
//         fmt.Println(r.Scope)
//         fmt.Println(r.Application.UID)
//         fmt.Println(r.CreatedAt.Time())
//         fmt.Println(r.User.ID)
//         fmt.Println(r.RateLimitInfo)
//     }
package goesa
