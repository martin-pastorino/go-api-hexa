package errors

type GenericError struct {
	Message string
	Code    int
}

func (e *GenericError) Error() string {
	return e.Message
}

func NewGenericError(message string) *GenericError {
return &GenericError{
		Message: message,
		Code:    500,
	}
}
