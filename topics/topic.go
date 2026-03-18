package topics

import nsq_client "github.com/RandySteven/go-kopi/pkg/nsq"

type Topics struct {
}

func NewTopics(nsq nsq_client.Nsq) *Topics {
	return &Topics{}
}
