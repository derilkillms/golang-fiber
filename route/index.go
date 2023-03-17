package route

import (
	"golang-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", "./public/asset")

	r.Get("/user", handler.UserHandlerRead)
	r.Get("/users", handler.UserHandlerGetAll)
	r.Get("/raws", handler.UserHandlerRaw)
	r.Post("/user/create", handler.UserHandlerCreate)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
