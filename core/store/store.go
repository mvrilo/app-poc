package store

import (
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto"
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
func Register(db *database.Database, grpcServer *grpc.Server, grpcClient grpc.ClientConnInterface) {
	service := New(db).Service
	proto.RegisterStoreServiceServer(grpcServer.Server, service)
	proto.NewStoreServiceClient(grpcClient)
}
