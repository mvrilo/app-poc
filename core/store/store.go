package store

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Store struct {
	*Service
}

func (s *Store) Register(ctx context.Context, db *database.Database, gs *grpc.Server, gc *grpc.Client) error {
	repo := &Repository{db}
	s.Service = &Service{repo}

	proto.RegisterStoreServiceServer(gs.Server, s.Service)
	proto.NewStoreServiceClient(gc)

	return proto.RegisterStoreServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
