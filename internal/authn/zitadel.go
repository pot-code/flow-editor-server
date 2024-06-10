package authn

import (
	"context"
	"flow-editor-server/internal/config"
	"fmt"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

func NewZitadelClient(config *config.HttpConfig) (*authorization.Authorizer[*oauth.IntrospectionContext], error) {
	z, err := authorization.New(
		context.Background(),
		zitadel.New(config.ZitadelDomain, zitadel.WithInsecure(config.ZitadelPort)),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to zitadel: %w", err)
	}
	return z, err
}
