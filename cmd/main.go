package main

import (
	"balancer/internal/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const port = 6080
const host = "localhost"
// const apiPrefix = "/api/v1"
// const apiGroup = "/group"
// const apiContact = "/contact"


func main() {
	svc := service.New()
	api := fiber.New()


	api.Get("/", svc.RootHandler)
	// api.Get(apiPrefix+apiGroup, svc.GroupHandler)
	// api.Get(apiPrefix+apiContact, svc.ContactHandler)

	if err := api.Listen(host + ":" + strconv.Itoa(port)); err != nil {
		log.Fatalln(err)
	}
}