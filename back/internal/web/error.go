package web

type Error struct {
	Code int   `json:"status"`
	Err  error `json:"error"`
}

func NewError(code int, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

func (e *Error) Status() int {
	return e.Code
}
