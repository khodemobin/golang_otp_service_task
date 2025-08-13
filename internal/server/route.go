package server

import (
	"github.com/khodemobin/golang_boilerplate/internal/server/middleware"
)

func (r *Server) routing() {
	api := r.router.Group("/api")

	auth := api.Group("/auth/otp")
	auth.Post("/send", r.handler.AuthHandler.SendOTP)
	auth.Post("/verify", r.handler.AuthHandler.VerifyOTP)

	user := api.Group("/users")
	user.Use(middleware.Protected(r.app))
	user.Get("/", r.handler.UserHandler.Index)
	user.Get("/:id", r.handler.UserHandler.Get)
}
