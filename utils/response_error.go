package utils

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}
