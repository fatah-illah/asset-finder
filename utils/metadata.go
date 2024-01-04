package utils

import "strconv"

type Metadata struct {
	PageNo   int
	PageSize int
	SearchBy string
}

func (e *ResponseError) Error() string {
	return "HTTP " + strconv.Itoa(e.Status) + ": " + e.Message
}
