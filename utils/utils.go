package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func GetHeaders(headers interface{}, hostname string) context.Context {
	fmt.Println("#######################")
	fmt.Println("headers: ", headers.(http.Header))
	headersMap := make(map[string]string)
	for key, values := range headers.(http.Header) {
		headersMap[key] = values[0]
	}
	fmt.Println("headers: ", headersMap)

	ctx := context.Background()

	md := metadata.New(headersMap)
	md = metadata.New(map[string]string{"hostname": hostname})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}
