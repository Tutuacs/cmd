package pubsub

import (
	"github.com/Tutuacs/pkg/cache"
	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/logs"
)

type PubSubService struct {
	addr string
	*cache.RedisPubSub
}

func UsePubSubService() (*PubSubService, error) {
	// cache service mantain always the same connection alive
	pubsub, err := cache.UseRedisPubSub()
	if err != nil {
		return nil, err
	}

	conf := config.GetRedis()

	return &PubSubService{
		RedisPubSub: pubsub,
		addr:        conf.Addr,
	}, nil
}

func (p *PubSubService) Run() {

	// You can subscribe to your channels and define the handlers here
	p.Subscribe("/hello", cache.HandleHello)
	p.Subscribe("/unknown", cache.HandleUnknown)

	p.Listen()
	logs.OkLog("Listening redis pub/sub messages...")
}
