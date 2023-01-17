package clientgrpc

import (
	"context"
	"github.com/rhyme80/redis-client-library/infrastructure/servicegrpc"
	"github.com/rhyme80/redis-client-library/model"
)

type ClientGrpc struct {
	client servicegrpc.RedisServiceClient
}

func New(client servicegrpc.RedisServiceClient) ClientGrpc {
	return ClientGrpc{client: client}
}

func (s ClientGrpc) Get(ctx context.Context, key string) (string, error) {
	responseGet, err := s.client.Get(ctx, &servicegrpc.RequestGet{Key: key})
	if err != nil {
		return "", err
	}

	if responseGet.GetIsCacheKeyNotFound() {
		return "", model.ErrCacheKeyNotFound
	}

	return responseGet.GetValue(), nil
}

func (s ClientGrpc) Set(ctx context.Context, key, value string, expirationInSeconds int64) error {
	_, err := s.client.Set(ctx, &servicegrpc.RequestSet{
		Key:                 key,
		Value:               value,
		ExpirationInSeconds: expirationInSeconds,
	})
	if err != nil {
		return err
	}

	return nil
}
