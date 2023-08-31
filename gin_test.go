package blunder

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/rotisserie/eris"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Message struct {
	Id         *int   `json:"id" validate:"required,isNumber"`
	ExternalId string `json:"external_id" validate:"alpha"`
}

type Output struct {
	Message string `json:"message"`
}

type InvalidTypeMessage struct {
	Id string `json:"id"`
}

type MesssageWithoutValidation struct {
	Id int `json:"id"`
}

type WrappedLogger struct {
	zap *zap.Logger
}

func (l WrappedLogger) Error(err error) {
	l.zap.Error(err.Error())
}

func NewWrappedLogger(zap *zap.Logger) WrappedLogger {
	return WrappedLogger{zap: zap}
}

type Handlers struct {
	blunder *Blunder
}

func NewHandlers(blunder *Blunder) *Handlers {
	return &Handlers{blunder: blunder}
}

func (h *Handlers) A(c *gin.Context) {

	err := eris.New("some error occured")
	if err != nil {
		h.blunder.GinAdd(c, err)
		return
	}

	c.JSON(http.StatusOK, Output{Message: "ok"})
}

func (h *Handlers) B(c *gin.Context) {

	message := &Message{}

	errors := h.blunder.BindJson(c.Request, message)
	if len(errors) > 0 {
		for _, err := range errors {
			h.blunder.GinAdd(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, Output{Message: message.ExternalId})
}

func (h *Handlers) C(c *gin.Context) {

	message := &MesssageWithoutValidation{}

	errors := h.blunder.BindJson(c.Request, message)
	if len(errors) > 0 {
		for _, err := range errors {
			h.blunder.GinAdd(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, Output{Message: "success"})
}

func (h *Handlers) D(c *gin.Context) {
	panic("panic error")
}

func longRunning(ctx context.Context) (int, error) {
	count := 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			count = count + i
			time.Sleep(2 * time.Second)
		}
	}
	return count, nil
}

func (h *Handlers) E(c *gin.Context) {

	ctx, cancelFunc := context.WithCancel(c.Request.Context())
	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()
	count, err := longRunning(ctx)
	if err != nil {
		h.blunder.GinAdd(c, err)
		return
	}

	c.JSON(http.StatusOK, Output{Message: strconv.Itoa(count)})
}

func (h *Handlers) F(c *gin.Context) {
	err := &RecordNotFoundError{}
	err.Name = "user"
	err.Identifier = "123"
	h.blunder.GinAdd(c, err)
}

func validateIsNumber(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func SetupServer() {
	logger, _ := zap.NewProduction()
	wrappedLogger := NewWrappedLogger(logger)

	blunder := NewRFC()

	err := blunder.RegisterCustomValidation("isNumber", validateIsNumber)
	if err != nil {
		panic(err)
	}

	trans, _ := blunder.uni.GetTranslator("en")
	err = blunder.RegisterCustomTranslation("isNumber", trans, func(ut ut.Translator) error {
		return ut.Add("isNumber", "{0} must be a number", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(fe.Tag(), fe.Field())
		return t
	})
	if err != nil {
		panic(err)
	}

	recordNotFoundErr := RecordNotFoundError{}

	blunder.AddCustomerError(recordNotFoundErr)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(gin.CustomRecovery(blunder.GinRecovery))
	r.Use(blunder.GinErrorHandler(wrappedLogger))
	r.NoMethod(blunder.GinNoMethod)
	r.NoRoute(blunder.GinNoRoute)

	handlers := NewHandlers(blunder)
	r.LoadHTMLFiles("blunder.html")
	r.GET("errors", blunder.Html)
	r.GET("/a", handlers.A)
	r.POST("/b", handlers.B)
	r.POST("/c", handlers.C)
	r.GET("/d", handlers.D)
	r.GET("/e", handlers.E)
	r.GET("/f", handlers.F)

	errs, _ := errgroup.WithContext(context.Background())

	errs.Go(func() error {
		return r.Run()
	})
}

// func TestHtml(t *testing.T) {

// 	SetupServer()

// 	client := resty.New()

// 	resp, err := client.R().
// 		Get("http://localhost:8080/errors")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.IsError() {
// 		t.Fail()
// 	}

// 	assert.Equal(t, http.StatusOK, resp.StatusCode(), "status missmatch")

// 	assert.Equal(t, "", string(resp.Body()), "body missmatch")

// }

func TestBasicError(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/a")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Title: "The request failed due to an internal error.",
			},
		},
		Status: http.StatusInternalServerError,
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")

}

func TestPostSuccess(t *testing.T) {

	SetupServer()

	client := resty.New()

	id := 123

	message := &Message{
		Id:         &id,
		ExternalId: "asdf",
	}

	resp, err := client.R().
		SetBody(message).
		Post("http://localhost:8080/b")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsError() {
		t.Fail()
	}

	expectedOutput := Output{
		Message: "asdf",
	}

	expectedBytes, err := json.Marshal(expectedOutput)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPostWithoutBody(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		Post("http://localhost:8080/b")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "Id must have a value!",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.Id",
				Expression: "required",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "ExternalId can only contain alphabetic characters",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.ExternalId",
				Expression: "alpha",
				Argument:   "string",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPostWithoutBodyFr(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept-Language", "fr").
		Post("http://localhost:8080/b")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "Id must have a value!",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.Id",
				Expression: "required",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "ExternalId can only contain alphabetic characters",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.ExternalId",
				Expression: "alpha",
				Argument:   "string",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPostInvalidTypeParameter(t *testing.T) {

	SetupServer()

	client := resty.New()

	message := InvalidTypeMessage{
		Id: "asd",
	}

	resp, err := client.R().
		SetBody(message).
		Post("http://localhost:8080/b")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "ExternalId can only contain alphabetic characters",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.ExternalId",
				Expression: "alpha",
				Argument:   "string",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPostRequiredParameter(t *testing.T) {

	SetupServer()

	client := resty.New()

	message := Message{}

	resp, err := client.R().
		SetBody(message).
		Post("http://localhost:8080/b")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "Id must have a value!",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.Id",
				Expression: "required",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "ExternalId can only contain alphabetic characters",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Placement:  "field",
				Field:      "Message.ExternalId",
				Expression: "alpha",
				Argument:   "string",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPostInvalidJson(t *testing.T) {

	SetupServer()

	client := resty.New()

	message := InvalidTypeMessage{
		Id: "asd",
	}

	resp, err := client.R().
		SetBody(message).
		Post("http://localhost:8080/c")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:       "https://example.com/problems",
				Title:      "INVALID_PARAMETER",
				Detail:     "json: cannot unmarshal string into Go struct field MesssageWithoutValidation.id of type int",
				ReasonCode: 150,
				Reason:     "The request failed because it contained an invalid parameter or parameter value.",
				Action:     "Please correct the request as per the error description/details provided in the error response.",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestPanic(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/d")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Title: "The request failed due to an internal error.",
			},
		},
		Status: http.StatusInternalServerError,
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestContextTimeout(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/e")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Title: "The request failed due to an internal error.",
			},
		},
		Status: http.StatusInternalServerError,
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}

func TestCustomerError(t *testing.T) {

	SetupServer()

	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/f")

	if err != nil {
		t.Fatal(err)
	}

	if resp.IsSuccess() {
		t.Fail()
	}

	expectedErrorResponse := HTTPErrorResponse{
		Errors: []HTTPError{
			{
				Type:  "https://example.com/problems",
				Title: "Record user with 123 identifier not found",
			},
		},
	}

	expectedBytes, err := json.Marshal(expectedErrorResponse)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "status missmatch")

	assert.Equal(t, string(expectedBytes), string(resp.Body()), "body missmatch")
}
