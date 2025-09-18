package apiresp

import "net/http"

type ErrDetail struct {
	Code        errCode `json:"code"`
	Description string  `json:"description"`
}

type Pagination struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	TotalPages int32 `json:"totalPages"`
}

type Response struct {
	OK         bool        `json:"ok"`
	StatusCode int         `json:"statusCode"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []ErrDetail `json:"errors,omitempty"`
}

func Success(data any) Response {
	return Response{
		OK:         true,
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func FatalError() Response {
	return Response{
		OK:         false,
		StatusCode: http.StatusInternalServerError,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeFatal,
				Description: "Internal server error",
			},
		},
	}
}

func BadRequestError() Response {
	return Response{
		OK:         false,
		StatusCode: http.StatusBadRequest,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeBadRequest,
				Description: "Bad request",
			},
		},
	}
}

func UnauthorizedError() Response {
	return Response{
		OK:         false,
		StatusCode: http.StatusUnauthorized,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnauthorized,
				Description: "Unauthorized",
			},
		},
	}
}

func NotAuthenticatedError() Response {
	return Response{
		OK:         false,
		StatusCode: http.StatusForbidden,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeNotAuthenticated,
				Description: "Not Authenticated",
			},
		},
	}
}

func CommonError(err error) Response {
	return Response{
		OK:         false,
		StatusCode: http.StatusInternalServerError,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnknown,
				Description: err.Error(),
			},
		},
	}
}
