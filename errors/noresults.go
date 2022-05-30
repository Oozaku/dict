package errors

import "fmt"

type NoResults struct {
	ErrorMessage string
	ErrorCode    int
}

func (err NoResults) Error() string {
	return fmt.Sprintf("%d - %s", err.ErrorCode, err.ErrorMessage)
}
