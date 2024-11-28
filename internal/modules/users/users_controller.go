package users

import (
	"bulletproof-go/internal/core/context"
	"bulletproof-go/internal/core/router"
	"net/http"
)

type UsersController struct {
	Path    string
	Service UsersService
}

func NewUsersController(
	service *UsersService,
	path string,
) *UsersController {
	return &UsersController{
		Service: *service,
		Path:    path,
	}
}

func (uc *UsersController) RegisterRoutes(r *router.Router) {
	tr := r.SubRouter(uc.Path)

	tr.Get("", func(ctx *context.Context) {
		ctx.JSON(http.StatusOK, "list of users")
	})

	tr.Post("", func(ctx *context.Context) {
		var body interface{}
		ctx.Body(&body)
		ctx.JSON(200, body)
	})

	tr.Delete("/{id}", func(ctx *context.Context) {
		id, _ := ctx.GetParam("id")
		ctx.Error(400, id)
	})
}
