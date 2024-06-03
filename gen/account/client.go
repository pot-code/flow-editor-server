// Code generated by goa v3.16.2, DO NOT EDIT.
//
// account client
//
// Command:
// $ goa gen flow-editor-server/design

package account

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "account" service client.
type Client struct {
	GetAccountEndpoint goa.Endpoint
}

// NewClient initializes a "account" service client given the endpoints.
func NewClient(getAccount goa.Endpoint) *Client {
	return &Client{
		GetAccountEndpoint: getAccount,
	}
}

// GetAccount calls the "getAccount" endpoint of the "account" service.
func (c *Client) GetAccount(ctx context.Context) (res *AccountOutput, err error) {
	var ires any
	ires, err = c.GetAccountEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*AccountOutput), nil
}