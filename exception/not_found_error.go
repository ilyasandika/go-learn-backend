package exception

type NotFoundError struct {
	Message string
}

func (error *NotFoundError) Error() string {
	return error.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message}
}
