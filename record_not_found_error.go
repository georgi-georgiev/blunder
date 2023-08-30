package blunder

import (
	"fmt"
	"net/http"
)

type RecordNotFoundError struct {
	Name       string
	Identifier string
}

func (e RecordNotFoundError) Code() int {
	return http.StatusBadRequest
}

func (e RecordNotFoundError) ShouldAbort() bool {
	return true
}

func (e RecordNotFoundError) ToHTPPError() HTTPError {
	return HTTPError{
		Title: e.Error(),
	}
}

func (e RecordNotFoundError) WithTitle() string {
	return e.Error()
}

func (e RecordNotFoundError) Recovarable() bool {
	return false
}

func (e RecordNotFoundError) Error() string {
	return fmt.Sprintf("Record %s with %s identifier not found", e.Name, e.Identifier)
}
