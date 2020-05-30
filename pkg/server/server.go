package server

import (
	"context"
	"log"

	"github.com/mvrilo/storepoc/pkg/config"
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/http"
)

type Server struct {
	Ctx context.Context

	Database   *database.Database
	HttpServer *http.Server
	GrpcServer *grpc.Server

	quit chan struct{}
}

type RegisterableService interface {
	Register(s *Server) error
}

func New() (*Server, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	grpcServer, err := grpc.NewServer()
	if err != nil {
		return nil, err
	}

	httpServer := http.NewServer()
	httpServer.AddGrpcGateway("/api", grpcServer.GatewayMux)

	s := &Server{
		Database:   db,
		HttpServer: httpServer,
		GrpcServer: grpcServer,
		Ctx:        context.Background(),
	}

	return s, err
}

func (s *Server) Load(services ...RegisterableService) error {
	for _, service := range services {
		err := service.Register(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) Start() {
	log.Println("Running migrations")

	log.Println("Starting grpc server at", config.GrpcAddress())
	go s.GrpcServer.Start()

	log.Println("Starting http server at", config.HttpAddress())
	go s.HttpServer.Start()

	<-s.quit
}

func (s *Server) Stop() {
	s.GrpcServer.Close()
	s.HttpServer.Close()
	s.Database.Close()
	s.quit <- struct{}{}
}
