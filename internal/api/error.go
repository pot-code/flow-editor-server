package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

func ErrorHandler(ctx context.Context, w http.ResponseWriter, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return JsonResponse(http.StatusNotFound, w, map[string]any{
			"message": err.Error(),
		})
	}

	switch te := err.(type) {
	case validation.Errors:
		return JsonResponse(http.StatusBadRequest, w, map[string]any{
			"message": te.Error(),
		})
	default:
		return JsonResponse(http.StatusInternalServerError, w, map[string]any{
			"message": te.Error(),
		})
	}
}

func JsonResponse(code int, w http.ResponseWriter, data any) error {
	w.WriteHeader(code)
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
