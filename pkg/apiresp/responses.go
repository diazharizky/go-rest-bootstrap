package apiresp

type ErrDetail struct {
	Code        errCode `json:"code"`
	Description string  `json:"description"`
}

type Pagination struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	TotalPages int32 `json:"total_pages"`
}

type Response struct {
	OK         bool        `json:"ok"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Errors     []ErrDetail `json:"errors,omitempty"`
}

func Success(data interface{}) Response {
	return Response{
		OK:   true,
		Data: data,
	}
}

func FatalError() Response {
	return Response{
		OK: false,
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
		OK: false,
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
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnauthorized,
				Description: "Unauthorized",
			},
		},
	}
}

func CommonError(err error) Response {
	return Response{
		OK: false,
		Errors: []ErrDetail{
			{
				Code:        ErrCodeUnknown,
				Description: err.Error(),
			},
		},
	}
}
