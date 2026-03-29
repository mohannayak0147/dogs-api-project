package http

import (
	"errors"
	"fmt"
)

type HTTPError struct {
	err  error
	code int
}

// NewError creates an NewError error object.
// err can be a string or of type error.
// wrapMsg can be used to provide extra context to the error.x
func NewError(err interface{}, statusCode int, wrapMsg ...string) HTTPError {
	var e error
	switch err := err.(type) {
	case nil:
		e = nil
	case error:
		e = err
	case string:
		e = errors.New(err)
	default:
		panic("err must be an error or string")
	}

	if len(wrapMsg) == 0 {
		return HTTPError{e, statusCode}
	}
	return HTTPError{fmt.Errorf("%s: %s", wrapMsg[0], e), statusCode}
}

func (h HTTPError) StatusCode() int { return h.code }
func (h HTTPError) Unwrap() error   { return h.err }
func (h HTTPError) Error() string {
	if h.err == nil {
		return "<nil>"
	}
	return h.err.Error()
}
