package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func Headers(ctx context.Context, hostname string) context.Context {
	md := metadata.New(map[string]string{
		"hostname":         hostname,
		"application-code": fmt.Sprintf("%v", ctx.Value("Application-Code")),
		"request-id":       fmt.Sprintf("%v", ctx.Value("Request-Id")),
		"accept":           fmt.Sprintf("%v", ctx.Value("Accept")),
		//"accept-encoding":  fmt.Sprintf("%v", ctx.Value("Accept-Encoding")),
		//"connection":       fmt.Sprintf("%v", ctx.Value("Connection")),
		//"user-agent":       fmt.Sprintf("%v", ctx.Value("User-Agent")),
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}
