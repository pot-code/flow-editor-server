package authn

import (
	"context"
	"net/http"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/zerolog/log"
)

type keySetCache interface {
	GetKeySet() jwk.Set
	SetKeySet(set jwk.Set)
}

type defaultKeySetCache struct {
	keySet jwk.Set
	mu     sync.RWMutex
}

func (c *defaultKeySetCache) GetKeySet() jwk.Set {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.keySet
}

func (c *defaultKeySetCache) SetKeySet(set jwk.Set) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.keySet = set
}

var _ keySetCache = (*defaultKeySetCache)(nil)

func JwtValidation(issuer, jwkEndpoint, audience string) func(next http.Handler) http.Handler {
	client := resty.New()
	cache := new(defaultKeySetCache)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			set := cache.GetKeySet()
			if set == nil {
				res, err := client.R().Get(jwkEndpoint)
				if err != nil {
					log.Error().Err(err).Str("endpoint", jwkEndpoint).Msg("failed to get jwk keys")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				set, err = jwk.Parse(res.Body())
				if err != nil {
					log.Error().Err(err).Msg("failed to parse jwk keys")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				cache.SetKeySet(set)
			}

			t, err := jwt.ParseRequest(r, jwt.WithKeySet(set), jwt.WithIssuer(issuer), jwt.WithAudience(audience))
			if err != nil {
				log.Debug().Err(err).Msg("failed to parse jwt token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(injectToken(r.Context(), t)))
		})
	}
}

type contextKey int

const (
	jwtTokenKey contextKey = iota
)

func injectToken(ctx context.Context, t jwt.Token) context.Context {
	return context.WithValue(ctx, jwtTokenKey, t)
}

func TokenFromContext(ctx context.Context) jwt.Token {
	return ctx.Value(jwtTokenKey).(jwt.Token)
}
