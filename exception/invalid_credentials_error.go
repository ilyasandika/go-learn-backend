package exception

type InvalidCredentialsError struct {
	Message string
}

func (error *InvalidCredentialsError) Error() string {
	return error.Message
}

func NewInvalidCredentialsError(message string) *InvalidCredentialsError {
	return &InvalidCredentialsError{message}
}
