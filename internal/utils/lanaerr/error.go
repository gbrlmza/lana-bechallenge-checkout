package lanaerr

import "net/http"

type lanaError struct {
	Err        error
	StatusCode int
}

func (e lanaError) Error() string {
	return e.Err.Error()
}

func FromErr(err error) lanaError {
	if lErr, ok := err.(lanaError); ok {
		return lErr
	}

	return lanaError{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}

func New(err error, code int) lanaError {
	return lanaError{
		StatusCode: code,
		Err:        err,
	}
}

func Empty() lanaError {
	return lanaError{
		StatusCode: http.StatusInternalServerError,
	}
}

func (e lanaError) WithCode(code int) lanaError {
	e.StatusCode = code
	return e
}

func (e lanaError) WithErr(err error) lanaError {
	e.Err = err
	return e
}

func (e lanaError) GetStatusCode() int {
	return e.StatusCode
}

func (e lanaError) GetError() error {
	return e.Err
}
