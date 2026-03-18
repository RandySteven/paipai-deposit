package usecases

import (
	"github.com/RandySteven/go-kopi/caches"
	usecases_interfaces "github.com/RandySteven/go-kopi/interfaces/usecases"
	nsq_client "github.com/RandySteven/go-kopi/pkg/nsq"
	temporal_client "github.com/RandySteven/go-kopi/pkg/temporal"
	"github.com/RandySteven/go-kopi/repositories"
)

type Usecases struct {
	UserUsecase usecases_interfaces.UserUsecase
}

func NewUsecases(repositories *repositories.Repositories,
	redis *caches.Caches,
	nsq nsq_client.Nsq,
	temporal temporal_client.Temporal,
) *Usecases {
	return &Usecases{}
}
