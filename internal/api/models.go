package api

type ErrorResponse struct {
	Message string `json:"message"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PageResponse[T any] struct {
	Items  []T `json:"items"`
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
