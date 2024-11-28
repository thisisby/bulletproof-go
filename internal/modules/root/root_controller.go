package root

import "bulletproof-go/internal/core/router"

type RootController struct {
	Path    string
	Service RootService
}

func NewRootController(s *RootService, path string) *RootController {
	return &RootController{
		Path:    path,
		Service: *s,
	}
}

func (rc *RootController) RegisterRoutes(r *router.Router) {
	// r.Get("/hello", func(ctx *LessGo.Context) {
	// 	ctx.Send("Hello world")
	// })
}
