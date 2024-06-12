package errors

import (
	"fmt"
)

type ErrorCode struct {
	code int64
	msg  string
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", e.code, e.msg)
}

func Wrap(err error, msg string) error {
	if errCode, ok := err.(*ErrorCode); ok {
		errCode.msg = fmt.Sprintf("%v %v", msg, errCode.msg)
		return errCode
	}

	return New(msg)
}

func WrapF(err error, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return Wrap(err, msg)
}

func New(msg string, args ...interface{}) error {
	return &ErrorCode{
		msg: fmt.Sprintf(msg, args...),
	}
}

func NewCodeErr(code int64, msg string, args ...interface{}) error {
	return &ErrorCode{
		code: code,
		msg:  fmt.Sprintf(msg, args...),
	}
}

func CodeMsg(err error) (code int64, msg string) {
	if errCode, ok := err.(*ErrorCode); ok {
		return errCode.code, errCode.msg
	}

	return 0, err.Error()
}
