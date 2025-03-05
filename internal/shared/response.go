package shared

type ErrorResponse struct {
	Message string `json:"message"`
}

type IDResponse struct {
	ID *int `json:"id"`
}

func ErrorRes(errorString string) ErrorResponse {
	return ErrorResponse{
		Message: errorString,
	}
}

func IDRes(id *int) IDResponse {
	return IDResponse{
		ID: id,
	}
}
