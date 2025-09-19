package apiresp

import (
	"net/http"
)

type CommonError struct {
	Code        errCode `json:"code"`
	Description string  `json:"description"`
}

type InputError struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

type Pagination struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	TotalPages int32 `json:"totalPages"`
}

type Response[T any] struct {
	OK         bool        `json:"ok"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []T         `json:"errors,omitempty"`
}

func Ok(data any) (int, Response[CommonError]) {
	return http.StatusOK, Response[CommonError]{
		OK:   true,
		Data: data,
	}
}

func Created(data any) (int, Response[CommonError]) {
	return http.StatusCreated, Response[CommonError]{
		OK:   true,
		Data: data,
	}
}

func FatalError() (int, Response[CommonError]) {
	return http.StatusInternalServerError, Response[CommonError]{
		OK: false,
		Errors: []CommonError{
			{
				Code:        ErrCodeFatal,
				Description: "Internal server error",
			},
		},
	}
}

func BadRequestError() (int, Response[CommonError]) {
	return http.StatusBadRequest, Response[CommonError]{
		OK: false,
		Errors: []CommonError{
			{
				Code:        ErrCodeBadRequest,
				Description: "Bad request",
			},
		},
	}
}

func UnauthorizedError() (int, Response[CommonError]) {
	return http.StatusUnauthorized, Response[CommonError]{
		OK: false,
		Errors: []CommonError{
			{
				Code:        ErrCodeUnauthorized,
				Description: "Unauthorized",
			},
		},
	}
}

func NotAuthenticatedError() (int, Response[CommonError]) {
	return http.StatusForbidden, Response[CommonError]{
		OK: false,
		Errors: []CommonError{
			{
				Code:        ErrCodeNotAuthenticated,
				Description: "Not Authenticated",
			},
		},
	}
}

func UnknownError(err error) (int, Response[CommonError]) {
	return http.StatusInternalServerError, Response[CommonError]{
		OK: false,
		Errors: []CommonError{
			{
				Code:        ErrCodeUnknown,
				Description: err.Error(),
			},
		},
	}
}

func InputRequiredError(errors []InputError) (int, Response[InputError]) {
	return http.StatusBadRequest, Response[InputError]{
		OK:     false,
		Errors: errors,
	}
}
