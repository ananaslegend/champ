package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

func main() {
	s := fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			uri := ctx.Request.URI()
			path := strings.Trim(strings.ToLower(string(uri.Path())), "/")

			switch path {
			case "test":
				fmt.Println("hello from test")
				break
			default:
				fmt.Println("err")
			}
		},
	}

	err := s.ListenAndServe("0.0.0.0:5000")
	if err != nil {
		fmt.Println(err)
	}
}
