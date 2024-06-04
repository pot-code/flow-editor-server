package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
)

func NewValidator(trans ut.Translator) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	zh_translation.RegisterDefaultTranslations(v, trans)
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return v
}

func NewTranslator() ut.Translator {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	return trans
}
