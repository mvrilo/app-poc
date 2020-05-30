package health

import (
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/server"
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

func (h *Health) Register(s *server.Server) error {
	h.Service = &Service{}
	gs := s.GrpcServer
	proto.RegisterHealthServiceServer(gs.Server, h.Service)
	return proto.RegisterHealthServiceHandlerFromEndpoint(
		s.Ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
