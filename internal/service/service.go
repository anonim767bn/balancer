package service

import (
	"balancer/internal/models"
	"log"
	"strconv"

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

func (s *Service) ContactGetHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	str_id := ctx.Params("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	return ctx.JSON(models.Contact{ID: id})
}

func (s *Service) ContactListGetHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	return ctx.JSON([]models.Contact{})
}

func (s *Service) GroupGetHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	str_id := ctx.Params("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	return ctx.JSON(models.Group{ID: id})
}

func (s *Service) GroupListGetHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	return ctx.JSON([]models.Group{})
}

func (s *Service) ContactPostHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(models.Contact{})
}

func (s *Service) GroupPostHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(models.Group{})
}

func (s *Service) GroupDeleteHandler(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNoContent)
	return nil
}
func (s *Service) ContactDeleteHandler(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNoContent)
	return nil
}

func (s *Service) GroupPutHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	str_id := ctx.Params("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	return ctx.JSON(models.Group{ID: id})
}

func (s *Service) ContactPutHandler(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	str_id := ctx.Params("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	return ctx.JSON(models.Contact{ID: id})
}
