package goa

import "goa.design/goa/v3/http"

type HttpRoute interface {
	MountRoute(mux http.ResolverMuxer)
}
