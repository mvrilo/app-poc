package store

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Store struct {
	*Service
	Client proto.StoreServiceClient
}

func New(gc *grpc.Client) *Store {
	return &Store{
		Client: proto.NewStoreServiceClient(gc),
	}
}

func (s *Store) Register(ctx context.Context, db *database.Database, gs *grpc.Server) error {
	repo := &Repository{db}
	s.Service = &Service{repo}
	db.AutoMigrate(&proto.Store{})
	proto.RegisterStoreServiceServer(gs.Server, s.Service)
	return proto.RegisterStoreServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
