package store

import (
	"context"

	"github.com/mvrilo/storepoc/pkg/validator"
	"github.com/mvrilo/storepoc/proto/v1"
)

type Service struct {
	repo *Repository
}

func (s *Service) Find(ctx context.Context, in *proto.FindRequest) (*proto.Store, error) {
	data := struct {
		Name string `validate:"required"`
	}{
		in.Name,
	}

	if err := validator.ValidateGrpc(data); err != nil {
		return nil, err
	}

	return s.repo.Find(in)
}

func (s *Service) Create(ctx context.Context, in *proto.CreateRequest) (*proto.Store, error) {
	data := struct {
		Name string `validate:"required"`
		Uri  string `validate:"required,uri"`
	}{
		in.Name,
		in.Uri,
	}

	if err := validator.ValidateGrpc(data); err != nil {
		return nil, err
	}

	return s.repo.Create(in)
}

func (s *Service) List(ctx context.Context, in *proto.ListRequest) (*proto.Stores, error) {
	return s.repo.List(in)
}

func (s *Service) ChangeStatus(ctx context.Context, in *proto.ChangeStatusRequest) (*proto.Store, error) {
	data := struct {
		Id     string       `validate:"required,uuid"`
		Status proto.Status `validate:"required"`
	}{
		in.Id,
		in.Status,
	}

	if err := validator.ValidateGrpc(data); err != nil {
		return nil, err
	}

	return s.repo.ChangeStatus(in)
}
