package blunder

type DefaultError interface {
	error
	ToHTPPError() HTTPError
}
