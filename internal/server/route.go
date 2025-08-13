package server

func (r *Server) routing() {
	api := r.router.Group("/api")

	auth := api.Group("/auth/otp")
	auth.Post("/send", r.handler.AuthHandler.SendOTP)
	auth.Post("/verify", r.handler.AuthHandler.VerifyOTP)
}
