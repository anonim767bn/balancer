package main

import (
	"balancer/internal/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const port = 7777
const host = "0.0.0.0"

const apiPrefix = "/api/v1"

const apiGroup = "/group"
const apiContact = "/contact"

func main() {
	svc := service.New()
	api := fiber.New()

	api.Get("/", svc.RootHandler)
	// api.Get(apiPrefix+apiGroup, svc.GroupHandler)
	api.Get(apiPrefix+apiContact, svc.ContactListGetHandler)
	api.Get(apiPrefix+apiContact+"/:id", svc.ContactGetHandler)

	api.Get(apiPrefix+apiGroup, svc.GroupListGetHandler)
	api.Get(apiPrefix+apiGroup+"/:id", svc.GroupGetHandler)

	api.Post(apiPrefix+apiContact, svc.ContactPostHandler)
	api.Post(apiPrefix+apiGroup, svc.GroupPostHandler)

	api.Delete(apiPrefix+apiContact+"/:id", svc.ContactDeleteHandler)
	api.Delete(apiPrefix+apiGroup+"/:id", svc.GroupDeleteHandler)

	api.Put(apiPrefix+apiContact+"/:id", svc.ContactPutHandler)
	api.Put(apiPrefix+apiGroup+"/:id", svc.GroupPutHandler)

	if err := api.Listen(host + ":" + strconv.Itoa(port)); err != nil {
		log.Fatalln(err)
	}
}
