package route

import (
	"golang-fiber/config"
	"golang-fiber/handler"
	"golang-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)
	r.Get("/user", handler.UserHandlerRead)
	r.Get("/users", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/raws", handler.UserHandlerRaw)
	r.Post("/user/create", handler.UserHandlerCreate)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
