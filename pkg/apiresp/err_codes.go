package apiresp

type errCode int32

const (
	ErrCodeFatal errCode = 101 + iota
	ErrCodeBadRequest
	ErrCodeUnauthorized
	ErrCodeUnknown
)
