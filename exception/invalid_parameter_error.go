package exception

type InvalidParameterError struct {
	Message string
}

func (e *InvalidParameterError) Error() string {
	return e.Message
}

func NewInvalidParameter(message string) *InvalidParameterError {
	return &InvalidParameterError{message}
}
