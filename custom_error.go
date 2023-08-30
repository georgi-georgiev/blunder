package blunder

type CustomError interface {
	DefaultError
	ShouldAbort() bool
	Code() int
	WithTitle() string
	Recovarable() bool
	// WithDetail(detail string)
	// WithRecovery()
	// WithAction(action string)
}
