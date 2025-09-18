package apiresp

type errCode int16

const (
	ErrCodeFatal errCode = 101 + iota
	ErrCodeBadRequest
	ErrCodeUnauthorized
	ErrCodeUnknown
	ErrCodeNotAuthenticated
)
