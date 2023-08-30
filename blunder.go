package blunder

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Blunder struct {
	valildator     *validator.Validate
	uni            *ut.UniversalTranslator
	environment    Environment
	typeURI        *string
	domain         *string
	isTraceable    bool
	isIdentifiable bool
	isTimeable     bool
	isRecovarable  bool
	isTranslatable bool

	customerErrors []CustomError
}

func New() *Blunder {

	Generate()

	//default en
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	val := validator.New()
	en_translations.RegisterDefaultTranslations(val, trans)

	//override
	val.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(fe.Tag(), fe.Field())

		return t
	})

	return &Blunder{
		valildator:  val,
		uni:         uni,
		environment: Development,
	}
}

// Consumers MUST use the "type" string as the primary identifier for
// the problem type; the "title" string is advisory and included only
// for users who are not aware of the semantics of the URI and do not
// have the ability to discover them (e.g., offline log analysis).
// Consumers SHOULD NOT automatically dereference the type URI.

// The "status" member, if present, is only advisory; it conveys the
// HTTP status code used for the convenience of the consumer.
// Generators MUST use the same status code in the actual HTTP response,
// to assure that generic HTTP software that does not understand this
// format still behaves correctly.  See Section 5 for further caveats
// regarding its use.

// Consumers can use the status member to determine what the original
// status code used by the generator was, in cases where it has been
// changed (e.g., by an intermediary or cache), and when message bodies
// persist without HTTP information.  Generic HTTP software will still
// use the HTTP status code.

// The "detail" member, if present, ought to focus on helping the client
// correct the problem, rather than giving debugging information.

// Consumers SHOULD NOT parse the "detail" member for information;
// extensions are more suitable and less error-prone ways to obtain such
// information.

// Note that both "type" and "instance" accept relative URIs; this means
// that they must be resolved relative to the document's base URI, as
// per [RFC3986], Section 5.
func NewRFC() *Blunder {
	blunder := New()
	defaultTypeURI := "https://example.com/problems"
	blunder.typeURI = &defaultTypeURI
	blunder.isIdentifiable = true
	return blunder
}

func NewWithOptions(options BlunderOptions) *Blunder {
	blunder := New()
	blunder.environment = options.Environment
	blunder.typeURI = &options.TypeURI
	blunder.domain = &options.Domain
	blunder.isTraceable = options.IsTraceable
	blunder.isIdentifiable = options.IsIdentifiable
	blunder.isTimeable = options.IsTimeable
	blunder.isRecovarable = options.IsRecovarable

	return blunder
}

func (b *Blunder) SetEnvironment(environment Environment) {
	b.environment = environment
}

func (b *Blunder) AddCustomerError(customerError CustomError) {
	b.customerErrors = append(b.customerErrors, customerError)
}

func (b *Blunder) RegisterCustomValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return b.valildator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func (b *Blunder) RegisterCustomTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) error {
	return b.valildator.RegisterTranslation(tag, trans, registerFn, translationFn)
}
