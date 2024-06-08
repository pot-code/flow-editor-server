package authz

type UnAuthorizedError struct {
	err error
}

func NewUnAuthorizedError(err error) *UnAuthorizedError {
	return &UnAuthorizedError{err: err}
}

func (e *UnAuthorizedError) Error() string {
	return e.err.Error()
}
