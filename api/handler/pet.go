package handler

import (
	"fiberent/api/presenter"
	"fiberent/entity"
	"fiberent/usecase/pet"

	"github.com/gofiber/fiber/v2"
)

func NewPetHandler(app fiber.Router, service pet.UseCase) {
	app.Post("/", createPet(service))
	app.Get("/", listPets(service))
	app.Get("/:petId", getPet(service))
	app.Post("/:petId", updatePet(service))
	app.Delete("/:petId", deletePet(service))
}

func createPet(service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		var pet *entity.Pet
		err := c.BodyParser(&pet)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		pet, err = service.CreatePet(ctx, pet.Name, pet.Age)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func getPet(service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("petId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		pet, err := service.GetPet(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Pet Found",
			"data":    toJ,
		})
	}
}

func updatePet(service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("petId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var pet *entity.Pet

		err = c.BodyParser(&pet)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		pet.ID = id

		pet, err = service.UpdatePet(ctx, pet)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Pet{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func deletePet(service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("petId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "Bad Id Format",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		err = service.DeletePet(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error deleting pet",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "pet deleted successfully",
			"error":  nil,
		})
	}
}

func listPets(service pet.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		users, err := service.ListPets(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := make([]presenter.Pet, len(users))

		for i, pet := range users {
			toJ[i] = presenter.Pet{
				ID:   pet.ID,
				Name: pet.Name,
				Age:  pet.Age,
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Pets Found",
			"data":    toJ,
		})
	}
}
