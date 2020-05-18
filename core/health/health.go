package health

import (
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/proto"
)

type Health struct {
	*Service
}

func New(_ *database.Database) *Health {
	return &Health{&Service{}}
}

// Register reisters a grpc client and server routes
func Register(db *database.Database, grpcServer *grpc.Server, grpcClient grpc.ClientConnInterface) {
	service := New(db).Service
	proto.RegisterHealthServiceServer(grpcServer.Server, service)
	proto.NewHealthServiceClient(grpcClient)
}
