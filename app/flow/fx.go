//go:build !goverter

package flow

import (
	"flow-editor-server/gen/flow"
	"flow-editor-server/gen/http/flow/server"
	"flow-editor-server/internal/goa"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	zw "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"go.uber.org/fx"
	"goa.design/goa/v3/http"
)

var Module = fx.Module(
	"flow",
	fx.Provide(
		fx.Annotate(NewService, fx.As(new(flow.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(
		s flow.Service,
		mux http.ResolverMuxer,
		zitadel *authorization.Authorizer[*oauth.IntrospectionContext],
		validator *validator.Validate,
		trans ut.Translator,
	) {
		endpoints := flow.NewEndpoints(s)
		endpoints.Use(goa.ValidatePayload(validator, trans))
		srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.ErrorFormatter)
		srv.Use(zw.New(zitadel).RequireAuthorization())
		server.Mount(mux, srv)
	}),
)
