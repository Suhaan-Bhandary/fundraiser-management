package internal_errors

type InvalidCredentialError struct {
	Message string
}

func (err InvalidCredentialError) Error() string {
	return err.Message
}

type NotFoundError struct {
	Message string
}

func (err NotFoundError) Error() string {
	return err.Message
}

type BadRequest struct {
	Message string
}

func (err BadRequest) Error() string {
	return err.Message
}
