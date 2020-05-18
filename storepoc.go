package storepoc

import (
	"log"

	"github.com/mvrilo/storepoc/pkg/config"
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"

	"github.com/mvrilo/storepoc/core/store"

	"github.com/mvrilo/storepoc/proto"
)

type Storepoc struct {
	*database.Database
	GrpcServer *grpc.Server
	GrpcClient *grpc.Client
}

func New() (*Storepoc, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	grpcServer, err := grpc.NewServer(config.GrpcAddress())
	if err != nil {
		return nil, err
	}

	grpcClient, err := grpc.NewClient(config.GrpcAddress())
	if err != nil {
		return nil, err
	}

	s := &Storepoc{
		Database:   db,
		GrpcServer: grpcServer,
		GrpcClient: grpcClient,
	}

	s.Router()
	return s, nil
}

func (s *Storepoc) Router() {
	store.Register(s.Database, s.GrpcServer, s.GrpcClient)
}

func (s *Storepoc) Start() {
	log.Println("Running migrations")
	s.Database.AutoMigrate(
		&proto.Store{},
	)
	log.Println("Starting grpc server at", config.GrpcAddress())
	s.GrpcServer.Start()
}

func (s *Storepoc) Stop() {
	s.Database.Close()
	s.GrpcServer.Close()
}
