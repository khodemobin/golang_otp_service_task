package server

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/khodemobin/golang_boilerplate/internal/app"
	"github.com/khodemobin/golang_boilerplate/internal/server/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app     *app.App
	router  *fiber.App
	handler *handler.Handler
}

func New(app *app.App) *Server {
	return &Server{
		app: app,
		router: fiber.New(fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				app.Log.Error(err)
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			},
		}),
		handler: handler.NewHandler(app.Log, app.Service),
	}
}

func (r *Server) Start() error {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "../..")

	r.router.Use(fiberLogger.New())
	r.router.Use(recover.New(), compress.New())
	r.router.Use(cors.New())
	r.router.Use(swagger.New(swagger.Config{
		FilePath: path + "/docs/swagger.json",
		Path:     "swagger",
	}))

	r.routing()
	return r.router.Listen(fmt.Sprintf(":%d", r.app.Config.App.Port))
}

func (r *Server) Shutdown() error {
	return r.router.Shutdown()
}
