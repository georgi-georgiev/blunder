package blunder

import (
	"net/http"
)

func (b *Blunder) HandleErrors(r *http.Request, errs []error) (int, HTTPErrorResponse, bool) {

	if len(errs) == 0 {
		return 0, HTTPErrorResponse{}, false
	}

	var status int
	var response HTTPErrorResponse

	for _, err := range errs {
		switch e := err.(type) {
		case CustomError:
			status = e.Code()
			httpErr := e.ToHTPPError()
			httpErr = b.Enrich(httpErr)
			response.Recoverable = e.Recovarable()
			response.Errors = append(response.Errors, httpErr)
			if e.ShouldAbort() {
				break
			}
		case validationError:
			status = http.StatusBadRequest
			httpErr := e.ToHTPPError()
			httpErr = b.Enrich(httpErr)
			response.Errors = append(response.Errors, httpErr)
			response.Recoverable = false
		default:
			return http.StatusInternalServerError, InternalServerError(), true
		}
	}

	if b.domain != nil {
		response.Domain = *b.domain
	}

	if b.isTimeable {
		response = response.WithTimestamp()
	}

	if b.isTraceable {
		response = response.WithTrace()
	}

	if b.isIdentifiable {
		response.Instance = r.URL.RawPath
	}

	if b.isTranslatable {
		language := r.Header.Get("Accept-Language")
		response.WithLanguage(language)
	}

	return status, response, true
}
