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

func (h *Health) Register(ctx context.Context, db *database.Database, gs *grpc.Server, gc *grpc.Client) error {
	h.Service = &Service{}

	proto.NewHealthServiceClient(gc)
	proto.RegisterHealthServiceServer(gs.Server, h.Service)

	return proto.RegisterHealthServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
