package main

import (
	"context"
	_ "fiberent/api/docs"
	"fmt"
	"log"

	"github.com/gofiber/swagger"

	"fiberent/api/handler"
	"fiberent/ent"
	"fiberent/infrastructure/ent/datastore"
	"fiberent/infrastructure/ent/repository"
	"fiberent/pkg/config"
	"fiberent/usecase/pet"
	"fiberent/usecase/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/lib/pq"
)

// @title			Fiber API
// @version		1.0
// @description	This is swagger for Fiber
// @host			localhost:3030
// @BasePath		/
func main() {

	config.LoadEnvironmentFile(".env")
	client := newDBClient()
	defer client.Close()

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "Fiber Ent Clean Architecture",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	// Create repositories.
	userRepository := repository.NewUserRepoEnt(client)
	petRepository := repository.NewPetRepoEnt(client)

	// Create all of our services.
	userService := user.NewService(userRepository)
	petService := pet.NewService(petRepository)

	api := app.Group("/api")

	// Prepare our endpoints for the API.
	handler.NewUserHandler(api.Group("/v1/users"), context.Background(), userService)
	handler.NewPetHandler(api.Group("/v1/pets"), context.Background(), petService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 3000.s
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening Posgres client: %v", err)
	}

	return client
}
