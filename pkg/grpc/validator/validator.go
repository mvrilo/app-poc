package validator

import (
	"github.com/mvrilo/storepoc/pkg/grpc"
	"github.com/mvrilo/storepoc/pkg/validator"
)

func Validate(data interface{}) error {
	if err := validator.Validate(data); err != nil {
		return grpc.Invalid(err)
	}
	return nil
}
