package flow

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateFlowInput struct {
	Title string  `json:"title" validate:"required"`
	Nodes *string `json:"nodes"`
	Edges *string `json:"edges"`
	Owner string  `json:"owner"`
}

func (c *CreateFlowInput) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 32)),
		validation.Field(&c.Owner, validation.Required),
	)
}

type FlowListObjectOutput struct {
	Id        int    `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
}

type UpdateFlowInput struct {
	Id    int     `json:"id" validate:"required"`
	Title string  `json:"title" validate:"required"`
	Nodes *string `json:"nodes"`
	Edges *string `json:"edges"`
	Owner string  `json:"owner"`
}

func (c *UpdateFlowInput) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 32)),
		validation.Field(&c.Owner, validation.Required),
	)
}

type FlowDetailOutput struct {
	Id        int     `json:"id" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Nodes     *string `json:"nodes"`
	Edges     *string `json:"edges"`
	CreatedAt string  `json:"created_at" validate:"required"`
}
