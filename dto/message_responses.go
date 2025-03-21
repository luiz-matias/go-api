package dto

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
}
