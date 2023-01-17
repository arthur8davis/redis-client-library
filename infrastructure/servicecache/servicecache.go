package servicecache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rhyme80/redis-client-library/infrastructure/clientgrpc"
	"github.com/rhyme80/redis-client-library/infrastructure/utils"
	"google.golang.org/grpc/metadata"
)

type ServiceCache struct {
	client      clientgrpc.ClientGrpc
	hoursExpire int
	hostname    string
}

func New(client clientgrpc.ClientGrpc, hoursExpire int, hostname string) ServiceCache {
	return ServiceCache{client: client, hoursExpire: hoursExpire, hostname: hostname}
}

func (s ServiceCache) SaveAll(key string, equivalentModel interface{}) error {
	bytes, err := json.Marshal(equivalentModel)
	if err != nil {
		return fmt.Errorf("cache.suggestedproducts.json.Marshal(): %w", err)
	}

	expireInSeconds := s.hoursExpire * 3600

	ctx := context.Background()
	md := metadata.New(map[string]string{"Request-Id": "req-123", "hostname": "hostname-321"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	if err := s.client.Set(ctx, key, string(bytes), int64(expireInSeconds)); err != nil {
		return err
	}
	return err
}

func (s ServiceCache) GetAll(ctx context.Context, key string) (interface{}, error) {
	ctx = utils.Headers(ctx, s.hostname)

	stringData, err := s.client.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var responses interface{}
	if err = json.Unmarshal([]byte(stringData), &responses); err != nil {
		return nil, fmt.Errorf("cache.suggestedproducts.json.Unmarshal(): %w", err)
	}

	return responses, nil
}
