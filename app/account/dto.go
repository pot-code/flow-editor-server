package account

type AccountOutput struct {
	Activated  bool `json:"activated" validate:"required"`
	Membership int  `json:"membership" validate:"required"`
}
