package authn

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/zerolog/log"
)

type keySetProvider struct {
	keySet   jwk.Set
	mu       sync.Mutex
	endpoint string
	rc       *resty.Client
}

func newKeySetProvider(endpoint string) *keySetProvider {
	return &keySetProvider{
		endpoint: endpoint,
		rc:       resty.New(),
	}
}

func (c *keySetProvider) getKeySet() (jwk.Set, error) {
	set := c.keySet
	if set != nil {
		return set, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.keySet != nil {
		return c.keySet, nil
	}

	log.Debug().Msgf("request jwk keys from %s", c.endpoint)
	res, err := c.rc.R().Get(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to request endpoint %s: %w", c.endpoint, err)
	}

	set, err = jwk.Parse(res.Body())
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwk keys: %w", err)
	}
	c.keySet = set
	return set, nil
}

func JwtValidation(jwkSource, issuer, audience string) func(next http.Handler) http.Handler {
	kp := newKeySetProvider(jwkSource)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			set, err := kp.getKeySet()
			if err != nil {
				log.Error().Err(err).Msg("failed to get jwk keys")
				w.WriteHeader(http.StatusInternalServerError)
				return
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
