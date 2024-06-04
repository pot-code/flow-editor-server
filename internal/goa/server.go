package goa

import "goa.design/goa/v3/http"

type HttpServer interface {
	MountHttpServer(mux http.ResolverMuxer)
}
