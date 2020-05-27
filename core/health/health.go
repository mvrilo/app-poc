package health

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Health struct {
	*Service
	Client proto.HealthServiceClient
}

func New(gc *grpc.Client) *Health {
	return &Health{
		Client: proto.NewHealthServiceClient(gc),
	}
}

func (h *Health) Register(ctx context.Context, db *database.Database, gs *grpc.Server) error {
	h.Service = &Service{}
	proto.RegisterHealthServiceServer(gs.Server, h.Service)
	return proto.RegisterHealthServiceHandlerFromEndpoint(
		ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
