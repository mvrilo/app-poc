package store

import (
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/server"
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

func (s *Store) Register(srv *server.Server) error {
	db := srv.Database
	repo := &Repository{db}
	s.Service = &Service{repo}
	db.AutoMigrate(&proto.Store{})

	gs := srv.GrpcServer
	proto.RegisterStoreServiceServer(gs.Server, s.Service)
	return proto.RegisterStoreServiceHandlerFromEndpoint(
		srv.Ctx,
		gs.GatewayMux,
		gs.GatewayAddr(),
		gs.GatewayOpts(),
	)
}
