package goa

import (
	"context"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	goa "goa.design/goa/v3/pkg"
)

func ValidatePayload(validate *validator.Validate, trans ut.Translator) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			rt := reflect.ValueOf(req)
			if rt.Kind() == reflect.Ptr && rt.Elem().Kind() == reflect.Struct {
				err := validate.Struct(req)
				if err != nil {
					if errs, ok := err.(validator.ValidationErrors); ok {
						return nil, goa.DecodePayloadError(errs[0].Translate(trans))
					}
					return nil, err
				}
			}
			return e(ctx, req)
		}
	}
}
