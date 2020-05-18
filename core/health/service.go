package health

import (
	"context"

	"github.com/mvrilo/storepoc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct{}

func (s *Service) Check(ctx context.Context, empty *emptypb.Empty) (*proto.Health, error) {
	return &proto.Health{Alive: true}, nil
}
