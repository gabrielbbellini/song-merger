package utils

type Exception struct {
	Message string
	Code    int
}

func (e Exception) Error() string {
	return e.Message
}
