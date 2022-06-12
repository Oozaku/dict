package errors

import "fmt"

type ReqError struct {
	Code    int
	Message string
}

func (err *ReqError) Error() string {
	return fmt.Sprintf("%d - %s", err.Code, err.Message)
}
