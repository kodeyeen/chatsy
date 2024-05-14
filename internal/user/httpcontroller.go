package user

type httpController struct {
	service service
}

func NewHTTPController(svc service) *httpController {
	return &httpController{
		service: svc,
	}
}
