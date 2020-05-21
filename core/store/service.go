package store

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/grpc/validator"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Service struct {
	repo *Repository
}

func (s *Service) Find(ctx context.Context, in *proto.FindRequest) (*proto.Store, error) {
	return s.repo.Find(in)
}

func (s *Service) Create(ctx context.Context, in *proto.CreateRequest) (*proto.Store, error) {

	if err := validator.Validate(in); err != nil {
		return nil, err
	}

	return s.repo.Create(in)
}

func (s *Service) List(ctx context.Context, in *proto.ListRequest) (*proto.Stores, error) {
	return s.repo.List(in)
}

func (s *Service) ChangeStatus(ctx context.Context, in *proto.ChangeStatusRequest) (*proto.Store, error) {
	return s.repo.ChangeStatus(in)
}
