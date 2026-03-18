package consumers

import (
	"context"
	"fmt"
	"log"

	"github.com/RandySteven/go-kopi/caches"
	nsq_client "github.com/RandySteven/go-kopi/pkg/nsq"
	"github.com/RandySteven/go-kopi/repositories"
	"github.com/RandySteven/go-kopi/topics"
)

type (
	Consumers struct {
		//DummyConsumer       consumer_interfaces.DummyConsumer
	}

	ConsumerFunc func(ctx context.Context) error

	RunConsumer map[string]ConsumerFunc

	Runners struct {
		nsq          nsq_client.Nsq
		ConsumerFunc []ConsumerFunc
		RunConsumers RunConsumer
	}
)

func InitRunner(nsq nsq_client.Nsq) *Runners {
	return &Runners{
		nsq:          nsq,
		RunConsumers: make(map[string]ConsumerFunc),
	}
}

func (r *Runners) RegisterConsumer(topic string, fun ConsumerFunc) *Runners {
	r.RunConsumers[topic] = fun
	return r
}

func (r *Runners) Run(ctx context.Context) error {
	errChan := make(chan error, len(r.RunConsumers))

	for topic, consumer := range r.RunConsumers {
		go func(topic string, consumer ConsumerFunc) {
			log.Println(`execute consumer `, consumer)
			err := r.nsq.RegisterConsumer(topic, func(msgCtx context.Context, key string) {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic in consumer %s: %v", topic, r)
					}
				}()

				if err := consumer(msgCtx); err != nil {
					log.Printf("Error in consumer %s: %v", topic, err)
				}
			})
			if err != nil {
				errChan <- fmt.Errorf("failed to register consumer for topic %s: %w", topic, err)
			}
		}(topic, consumer)
	}

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}

}

func NewConsumers(
	repo *repositories.Repositories,
	cache *caches.Caches,
	topics *topics.Topics,
) *Consumers {
	return &Consumers{}
}
