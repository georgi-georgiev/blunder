package blunder

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (b *Blunder) BindJson(r *http.Request, payload interface{}) []validationError {
	decodeErr := json.NewDecoder(r.Body).Decode(payload)

	err := b.valildator.Struct(payload)

	if err != nil {
		ve, ok := err.(validator.ValidationErrors)

		if ok {
			language := r.Header.Get("Accept-Language")
			return b.FromValidator(ve, language)
		}

		return FromError(err)
	}

	if decodeErr != nil {
		return FromError(decodeErr)
	}

	return nil
}
