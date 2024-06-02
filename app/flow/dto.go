package flow

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateFlowData struct {
	Title string  `json:"title"`
	Nodes *string `json:"nodes"`
	Edges *string `json:"edges"`
	Owner string  `json:"owner"`
}

func (c *CreateFlowData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 32)),
		validation.Field(&c.Owner, validation.Required),
	)
}

type FlowListItem struct {
	Id        int
	Title     string
	CreatedAt time.Time
}

type UpdateFlowData struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Nodes *string `json:"nodes"`
	Edges *string `json:"edges"`
	Owner string  `json:"owner"`
}

func (c *UpdateFlowData) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 32)),
		validation.Field(&c.Owner, validation.Required),
	)
}
