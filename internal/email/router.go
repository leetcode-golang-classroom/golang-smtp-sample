package email

import "net/http"

type Router struct {
	router  *http.ServeMux
	handler *Handler
}

func NewRouter(handler *Handler, router *http.ServeMux) *Router {
	return &Router{handler: handler, router: router}
}

func (r *Router) SetupRoute() {
	r.router.HandleFunc("POST /email", r.handler.EmailHandler)
	r.router.HandleFunc("POST /html_template_email", r.handler.HTMLTemplateEmailHandler)
}
