package http

import "net/http"

type ServeMux struct {
	*http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{}
}

type Server struct {
	*http.Server
}
