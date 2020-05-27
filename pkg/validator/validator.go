package validator

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/mvrilo/storepoc/pkg/grpc"
)

var (
	trans            ut.Translator
	uni              *ut.UniversalTranslator
	DefaultValidator *validator.Validate
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	DefaultValidator = validator.New()
	en_translations.RegisterDefaultTranslations(DefaultValidator, trans)
}

func Validate(data interface{}) (err error) {
	if err = DefaultValidator.Struct(data); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return errors.New(errs[0].Translate(trans))
		}
	}
	return
}

func ValidateGrpc(data interface{}) error {
	if err := Validate(data); err != nil {
		return grpc.Invalid(err)
	}
	return nil
}
