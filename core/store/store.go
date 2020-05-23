package store

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Store struct {
	*Repository
	*Service
}

func New(db *database.Database) *Store {
	repo := &Repository{db}
	service := &Service{repo}
	return &Store{repo, service}
}

// Register reisters a grpc client and server routes
func Register(ctx context.Context, db *database.Database, gs *grpc.Server, gc grpc.ClientConnInterface) error {
	service := New(db).Service
	proto.RegisterStoreServiceServer(gs.Server, service)
	proto.NewStoreServiceClient(gc)

	return proto.RegisterStoreServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
