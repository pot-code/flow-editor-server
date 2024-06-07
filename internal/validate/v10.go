package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
)

func NewValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "" {
			return fld.Name
		}

		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return fld.Name
		}
		return name
	})
	return v
}

func NewTranslator(v *validator.Validate) ut.Translator {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	zh_translation.RegisterDefaultTranslations(v, trans)
	return trans
}
