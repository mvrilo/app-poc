package health

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Health struct {
	*Service
}

func New(_ *database.Database) *Health {
	return &Health{&Service{}}
}

// Register reisters a grpc client and server routes
func Register(ctx context.Context, db *database.Database, gs *grpc.Server, gc grpc.ClientConnInterface) error {
	service := New(db).Service

	proto.NewHealthServiceClient(gc)
	proto.RegisterHealthServiceServer(gs.Server, service)

	return proto.RegisterHealthServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
