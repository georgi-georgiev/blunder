package blunder

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ctxKey struct{}

func (b *Blunder) Add(r *http.Request, err error) {
	ctx := r.Context()
	value := ctx.Value(ctxKey{})
	errors := make([]error, 0)
	if value != nil {
		errors, _ = value.([]error)
	}

	errors = append(errors, err)

	newCtx := context.WithValue(ctx, ctxKey{}, errors)
	_ = r.WithContext(newCtx)
}

func (b *Blunder) Get(r *http.Request) []error {
	ctx := r.Context()
	value := ctx.Value(ctxKey{})
	if value != nil {
		errors, ok := value.([]error)

		if ok {
			return errors
		}
	}

	return []error{}
}

func (b *Blunder) GinAdd(c *gin.Context, err error) {
	_ = c.Error(err)
}

func (b *Blunder) GinGet(c *gin.Context) []error {
	var errors []error

	for _, err := range c.Errors {
		errors = append(errors, err.Err)
	}

	return errors
}
