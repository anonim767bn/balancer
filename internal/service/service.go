package service

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) RootHandler(ctx *fiber.Ctx) error {
	log.Println("RootHandler have used")
	ctx.Set("Content-Type", "text/html")
	return ctx.SendString(`I'm alive, please use <a href="/api/v1/contact">contact api</a> or <a href="/api/v1/group">group api</a>`)
}
