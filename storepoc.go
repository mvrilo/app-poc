package storepoc

import (
	"context"
	"log"

	"github.com/mvrilo/storepoc/pkg/config"
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/http"

	"github.com/mvrilo/storepoc/core/health"
	"github.com/mvrilo/storepoc/core/store"

	"github.com/mvrilo/storepoc/proto/v1"
)

type Storepoc struct {
	*database.Database

	HttpServer *http.Server
	GrpcServer *grpc.Server
	GrpcClient *grpc.Client

	Ctx context.Context

	quit chan struct{}
}

type RegisterableService interface {
	Register(context.Context, *database.Database, *grpc.Server, *grpc.Client) error
}

func New() (*Storepoc, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	grpcClient, err := grpc.NewClient()
	if err != nil {
		return nil, err
	}

	grpcServer, err := grpc.NewServer()
	if err != nil {
		return nil, err
	}

	httpServer := http.NewServer()
	httpServer.AddGrpcGateway("/api", grpcServer.GatewayMux)

	s := &Storepoc{
		Database:   db,
		HttpServer: httpServer,
		GrpcServer: grpcServer,
		GrpcClient: grpcClient,
		Ctx:        context.Background(),
	}

	err = s.Load(
		&health.Health{},
		&store.Store{},
	)

	return s, err
}

func (s *Storepoc) Load(services ...RegisterableService) error {
	for _, service := range services {
		err := service.Register(s.Ctx, s.Database, s.GrpcServer, s.GrpcClient)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storepoc) Start() {
	log.Println("Running migrations")
	s.Database.AutoMigrate(
		&proto.Store{},
	)

	log.Println("Starting grpc server at", config.GrpcAddress())
	go s.GrpcServer.Start()

	log.Println("Starting http server at", config.HttpAddress())
	go s.HttpServer.Start()

	<-s.quit
}

func (s *Storepoc) Stop() {
	s.GrpcServer.Close()
	s.HttpServer.Close()
	s.Database.Close()
	s.quit <- struct{}{}
}
