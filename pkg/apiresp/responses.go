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

func Ok(data any) (int, Response) {
	return http.StatusOK, Response{
		OK:   true,
		Data: data,
	}
}

func FatalError() (int, Response) {
	return http.StatusInternalServerError, Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeFatal,
				Description: "Internal server error",
			},
		},
	}
}

func BadRequestError() (int, Response) {
	return http.StatusBadRequest, Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeBadRequest,
				Description: "Bad request",
			},
		},
	}
}

func UnauthorizedError() (int, Response) {
	return http.StatusUnauthorized, Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnauthorized,
				Description: "Unauthorized",
			},
		},
	}
}

func NotAuthenticatedError() (int, Response) {
	return http.StatusForbidden, Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeNotAuthenticated,
				Description: "Not Authenticated",
			},
		},
	}
}

func CommonError(err error) (int, Response) {
	return http.StatusInternalServerError, Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnknown,
				Description: err.Error(),
			},
		},
	}
}
