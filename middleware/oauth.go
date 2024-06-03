package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
)

func OauthInterceptor[T authorization.Ctx](authorizer *authorization.Authorizer[T], skipper func(c echo.Context) bool, options ...authorization.CheckOption) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			ctx, err := authorizer.CheckAuthorization(c.Request().Context(), c.Request().Header.Get(authorization.HeaderName), options...)
			if err != nil {
				if errors.Is(err, &authorization.UnauthorizedErr{}) {
					return c.JSON(http.StatusUnauthorized, echo.Map{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusForbidden, echo.Map{
					"error": err.Error(),
				})
			}
			c.SetRequest(c.Request().WithContext(authorization.WithAuthContext(c.Request().Context(), ctx)))
			return next(c)
		}
	}
}
