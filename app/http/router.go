package http

type Router struct {
	handlers []Controller
}

func (r *Router) AddHandler(controller Controller) {
	r.handlers = append(r.handlers, controller)
}

func (r *Router) Handle(request Request) Response {
	for _, controller := range r.handlers {
		res, ok := controller.Handle(request)
		if ok {
			return res
		}
	}
	return NoBody(NOT_FOUND)
}
