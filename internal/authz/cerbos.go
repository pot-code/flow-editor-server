package authz

import (
	"flow-editor-server/internal/config"
	"fmt"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

func NewCerbosClient(config *config.HttpConfig) (*cerbos.GRPCClient, error) {
	c, err := cerbos.New(config.CerobsAddr, cerbos.WithPlaintext())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cerbos: %w", err)
	}
	return c, err
}
