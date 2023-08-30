package blunder

import (
	"fmt"
	"net/http"
)

type ResourceAlreadyExistsError struct {
	Name       string
	Identifier string
}

func (e ResourceAlreadyExistsError) Code() int {
	return http.StatusBadRequest
}

func (e ResourceAlreadyExistsError) ShouldAbort() bool {
	return true
}

func (e ResourceAlreadyExistsError) ToHTPPError() HTTPError {
	return HTTPError{
		Title: e.Error(),
	}
}

func (e ResourceAlreadyExistsError) Error() string {
	return fmt.Sprintf("Resource %s with %s identifier already exists", e.Name, e.Identifier)
}
