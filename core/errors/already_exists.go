package errors

type AlreadyExists struct {
	Message string
	Code    int
}

func (e *AlreadyExists) Error() string {
	return e.Message
}

func NewAlreadyExists(message string) *AlreadyExists {
return &AlreadyExists{
		Message: message,
		Code:    409,
	}
}
