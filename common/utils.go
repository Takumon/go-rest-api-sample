package common

type CommonError struct {
	Error   error
	Code    int    `json:"code"`
	Message string `json:"message"`
}
