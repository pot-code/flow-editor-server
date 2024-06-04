package goa

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	ghttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	code    int    `json:"-"`
}

func (r *ErrorResponse) StatusCode() int { return r.code }

func ErrorFormatter(ctx context.Context, err error) ghttp.Statuser {
	if gerr, ok := err.(*goa.ServiceError); ok {
		return &ErrorResponse{
			ID:      gerr.ID,
			Message: gerr.Message,
			code:    ghttp.NewErrorResponse(ctx, gerr).StatusCode(),
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &ErrorResponse{ID: goa.NewErrorID(), Message: err.Error(), code: http.StatusNotFound}
	}
	log.Err(err).Msg("internal server error")
	return &ErrorResponse{ID: goa.NewErrorID(), Message: err.Error(), code: http.StatusInternalServerError}
}
