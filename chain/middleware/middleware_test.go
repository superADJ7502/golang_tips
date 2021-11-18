package middleware

import (
	"context"
	"fmt"
	"testing"
)

func annotate(s string) Middleware {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

func myEndpoint(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my endpoint!")
	return struct{}{}, nil
}

func Test_ExampleChain(t *testing.T) {
	e := Chain(
		annotate("first"),
		annotate("second"),
		annotate("third"),
	)(myEndpoint)

	if _, err := e(context.Background(), struct{}{}); err != nil {
		panic(err)
	}
}
