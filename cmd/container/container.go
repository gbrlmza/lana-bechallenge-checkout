package container

import (
	"context"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/domain/checkout"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/locker"
	"github.com/gbrlmza/lana-bechallenge-checkout/internal/repository/storage"
)

// TODO: Explain app wiring

func NewContainer(ctx context.Context) *checkout.Container {
	return &checkout.Container{
		Storage: storage.NewStorage(ctx),
		Locker:  locker.NewLocker(ctx),
	}
}
