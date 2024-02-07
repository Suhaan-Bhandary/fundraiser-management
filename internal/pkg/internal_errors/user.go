package internal_errors

type DuplicateKeyError struct {
	Message string
}

func (err DuplicateKeyError) Error() string {
	return err.Message
}
