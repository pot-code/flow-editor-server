package account

import (
	"context"
	"encoding/json"
	"flow-editor-server/gen/account"
	"net/http"

	goa "goa.design/goa/v3/pkg"
)

func Middleware(s account.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a, err := s.GetAccount(r.Context())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				b, _ := json.Marshal(map[string]any{
					"id":      goa.NewErrorID(),
					"message": err.Error(),
				})
				w.Write(b)
				return
			}
			next.ServeHTTP(w, r.WithContext(withContext(r.Context(), a)))
		})
	}
}

type accountKeyType string

const accountKey accountKeyType = "account"

func withContext(ctx context.Context, a *account.AccountInfo) context.Context {
	return context.WithValue(ctx, accountKey, a)
}

func Context(ctx context.Context) *account.AccountInfo {
	return ctx.Value(accountKey).(*account.AccountInfo)
}
