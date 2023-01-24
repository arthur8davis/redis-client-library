package clientgrpc

import (
	"context"
	"github.com/rhyme80/redis-client-library/model"
	"github.com/rhyme80/redis-client-library/servicegrpc"
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

func (s ClientGrpc) Del(ctx context.Context, key string) error {
	_, err := s.client.Del(ctx, &servicegrpc.RequestDel{
		Key: key,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s ClientGrpc) Expire(ctx context.Context, key string, expirationInSeconds int64) error {
	_, err := s.client.Expire(ctx, &servicegrpc.RequestExpire{
		Key:                 key,
		ExpirationInSeconds: expirationInSeconds,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s ClientGrpc) HGet(ctx context.Context, hash, key string) (string, error) {
	responseHGet, err := s.client.HGet(ctx, &servicegrpc.RequestHGet{
		Hash: hash,
		Key:  key,
	})
	if err != nil {
		return "", err
	}

	if responseHGet.GetIsCacheKeyNotFound() {
		return "", model.ErrCacheKeyNotFound
	}

	return responseHGet.GetValue(), nil
}

func (s ClientGrpc) HSet(ctx context.Context, hash, key, value string) error {
	_, err := s.client.HSet(ctx, &servicegrpc.RequestHSet{
		Hash:  hash,
		Key:   key,
		Value: value,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s ClientGrpc) HDel(ctx context.Context, hash, key string) error {
	_, err := s.client.HDel(ctx, &servicegrpc.RequestHDel{
		Hash: hash,
		Key:  key,
	})
	if err != nil {
		return err
	}

	return nil
}
