package blunder

import (
	"github.com/go-playground/validator/v10"
)

type validationError struct {
	Message    string
	Placement  string
	Field      string
	Expression string
	Argument   string
}

func (e validationError) Error() string {
	return "Validation error"
}

func (e validationError) ToHTPPError() HTTPError {
	reason := Reasons[INVALID_PARAMETER]

	httpErr := HTTPError{
		Type:       "https://example.com/problems/request-parameters-missing",
		Title:      INVALID_PARAMETER.String(),
		Detail:     e.Message,
		ReasonCode: int(INVALID_PARAMETER),
		Reason:     reason.Message,
		Action:     reason.ReasonGroup.Resolution,
		Placement:  e.Placement,
		Field:      e.Field,
		Expression: e.Expression,
		Argument:   e.Argument,
	}

	return httpErr
}

func (b *Blunder) FromValidator(errs validator.ValidationErrors, locale string) []validationError {
	validationErrors := make([]validationError, 0)

	for _, err := range errs {
		validationError := b.getValidationError(err, locale)
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

func FromError(err error) []validationError {
	return []validationError{
		{
			Message: err.Error(),
		},
	}
}

func (b *Blunder) getValidationError(fieldError validator.FieldError, locale string) validationError {
	field := fieldError.StructNamespace()
	tag := fieldError.Tag()

	trans, _ := b.uni.GetTranslator(locale)
	message := fieldError.Translate(trans)

	return validationError{
		Message:    message,
		Placement:  "field",
		Field:      field,
		Expression: tag,
		Argument:   fieldError.Type().Name(),
	}
}
