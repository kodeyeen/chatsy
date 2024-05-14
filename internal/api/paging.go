package api

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Page[T any] struct {
	Items  []*T `json:"items"`
	Count  int  `json:"count"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
}
