package middleware

import (
	"template_clean_architecture/utils/config"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	App *fiber.App
	Cfg *config.Config
}

func NewMiddleware(app *fiber.App, cfg *config.Config) *Middleware {
	return &Middleware{app, cfg}

}

func (m *Middleware) Register() {
	
}
