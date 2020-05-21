package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Invalid(err error) error {
	return status.Errorf(codes.InvalidArgument, err.Error())
}
