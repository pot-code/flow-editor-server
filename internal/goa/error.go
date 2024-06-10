package goa

import (
	"context"
	"errors"
	"flow-editor-server/internal/authz"
	"net/http"

	"github.com/rs/zerolog/log"
	ghttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"gorm.io/gorm"
)

type HttpErrorResponse struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	code    int    `json:"-"`
}

func (r *HttpErrorResponse) StatusCode() int { return r.code }

func HttpErrorFormatter(ctx context.Context, err error) ghttp.Statuser {
	if gerr, ok := err.(*goa.ServiceError); ok {
		return &HttpErrorResponse{
			ID:      gerr.ID,
			Message: gerr.Message,
			code:    ghttp.NewErrorResponse(ctx, gerr).StatusCode(),
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &HttpErrorResponse{ID: goa.NewErrorID(), Message: err.Error(), code: http.StatusNotFound}
	}
	if ue, ok := err.(*authz.UnAuthorizedError); ok {
		return &HttpErrorResponse{ID: goa.NewErrorID(), Message: ue.Error(), code: http.StatusForbidden}
	}
	log.Err(err).Msg("internal server error")
	return &HttpErrorResponse{ID: goa.NewErrorID(), Message: err.Error(), code: http.StatusInternalServerError}
}
