package blunder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func (b *Blunder) ErrorHandler(logger *zap.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)

			errors := b.Get(r)

			for _, err := range errors {
				logger.Error(err.Error())
				errors = append(errors, err)
			}

			statusCode, response, ok := b.HandleErrors(r, errors)
			if ok {
				w.Header().Set("Content-Type", "application/problem+json")
				w.WriteHeader(statusCode)
				_, err := w.Write(response.ToJson())
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func (b *Blunder) GinErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var errors []error

		for _, err := range c.Errors {
			logger.Error(err.Error())
			errors = append(errors, err.Err)
		}

		statusCode, response, ok := b.HandleErrors(c.Request, errors)
		if ok {
			c.Header("Content-Type", "application/problem+json")
			c.AbortWithStatusJSON(statusCode, response)
		}
	}
}

func (b *Blunder) GinNoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, NotFound())
}

func (b *Blunder) GinNoMethod(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, MethodNotAllowed())
}

func (b *Blunder) GinRecovery(c *gin.Context, err any) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError())
}

type ErrorCode struct {
	Status      string
	Title       string
	Description string
	Resolution  string
	Code        int
	Reason      string
	Message     string
	Tip         string
}

type Data struct {
	ErrorCodes []ErrorCode
}

func (b *Blunder) Html(c *gin.Context) {

	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "blunder.html", gin.H{})
}
