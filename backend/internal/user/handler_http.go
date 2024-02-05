package user

type handler struct {
	svc service
}

func NewHandler(s service) *handler {
	return &handler{
		svc: s,
	}
}
