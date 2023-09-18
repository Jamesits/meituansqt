package sqtapi

import (
	"fmt"
)

type Response[T any] struct {
	ApiError
	Data *T `json:"data"`
}

func (r *Response[T]) Error() error {
	if r.Status == 0 {
		return nil
	}

	return &r.ApiError
}

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func (re *ApiError) Error() string {
	return fmt.Sprintf("API error: %d %s", re.Status, re.Message)
}
