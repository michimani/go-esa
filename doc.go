// Package goesa is the unofficial esa API v1 SDK for the Go programming language.
//
// Getting started: To get information of OAuth access token.
//
//		go get github.com/michimani/go-esa/v2/gesa
//		go get github.com/michimani/go-esa/v2/esaapi/oauthtoken
//		go get github.com/michimani/go-esa/v2/esaapi/oauthtoken/types
//
//
//	    package main
//
//	    import (
//	        "context"
//	        "fmt"
//	        "os"
//
//	        "github.com/michimani/go-esa/v2/esaapi/oauthtoken"
//	        "github.com/michimani/go-esa/v2/esaapi/oauthtoken/types"
//	        "github.com/michimani/go-esa/v2/gesa"
//	    )
//
//	    func main() {
//	        token := "your-access-token"
//	        c, err := gesa.NewClient(&gesa.NewClientInput{
//	            AccessToken: token,
//	        })
//
//	        if err != nil {
//	            panic(err)
//	        }
//
//	        getOAuthTokenInfo(c)
//	    }
//
//	    func getOAuthTokenInfo(c *gesa.Client) {
//	        r, err := oauthtoken.GetOAuthTokenInfo(context.Background(), c, &types.GetOAuthTokenInfoInput{})
//	        if err != nil {
//	            fmt.Println(err)
//
//	            ge := err.(*gesa.GesaError)
//	            if ge.OnAPI {
//	                fmt.Println(ge.EsaAPIError.StatusCode)
//	                fmt.Println(ge.EsaAPIError.Status)
//	                fmt.Println(ge.EsaAPIError.Error)
//	                fmt.Println(ge.EsaAPIError.Message)
//
//	                if ge.EsaAPIError.RateLimitInfo != nil {
//	                    fmt.Println(ge.EsaAPIError.RateLimitInfo.Limit)
//	                    fmt.Println(ge.EsaAPIError.RateLimitInfo.Remaining)
//	                    fmt.Println(ge.EsaAPIError.RateLimitInfo.Reset)
//	                }
//	            }
//
//	            return
//	        }
//
//	        fmt.Printf("Response: %+v \n", r)
//
//	        fmt.Println(gesa.IntValue(r.ResourceOwnerID))
//	        fmt.Println(r.Scope)
//	        fmt.Println(r.Application.UID)
//	        fmt.Println(r.CreatedAt.Time())
//	        fmt.Println(r.User.ID)
//	        fmt.Println(r.RateLimitInfo)
//	    }
package goesa
