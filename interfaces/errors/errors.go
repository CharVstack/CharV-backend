package errors

import "github.com/pkg/errors"

type ErrorType uint

const (
	Unknown ErrorType = iota
	InvalidArgument
	NotFound
)

type typeGetter interface {
	Type() ErrorType
}

type customError struct {
	errorType     ErrorType
	originalError error
}

func (et ErrorType) New(message string) error {
	return customError{errorType: et, originalError: errors.New(message)}
}

func (et ErrorType) Wrap(err error, message string) error {
	return customError{errorType: et, originalError: errors.Wrap(err, message)}
}

func (e customError) Error() string {
	return e.originalError.Error()
}

func (e customError) Type() ErrorType {
	return e.errorType
}

func Wrap(err error, message string) error {
	wrapped := errors.Wrap(err, message)
	if ce, ok := err.(typeGetter); ok {
		return customError{errorType: ce.Type(), originalError: wrapped}
	}
	return customError{errorType: Unknown, originalError: wrapped}
}

func Cause(err error) error {
	return errors.Cause(err)
}

func GetType(err error) ErrorType {
	for {
		if e, ok := err.(typeGetter); ok {
			return e.Type()
		}
		break
	}
	return Unknown
}
