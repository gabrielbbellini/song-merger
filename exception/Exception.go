package utils

type Exception struct {
	Message string
	Code    int
}

func NewException(message string, code int) *Exception {
	return &Exception{
		Message: message,
		Code:    code,
	}
}

func (e Exception) Error() string {
	return e.Message
}

func (e Exception) ErrorCode() int {
	return e.Code
}
