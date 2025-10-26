package handler

import (
	"fiberent/api/presenter"
	"fiberent/entity"
	"fiberent/usecase/user"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, service user.UseCase) {
	app.Post("/", createUser(service))
	app.Get("/", listUsers(service))
	app.Get("/:userId", getUser(service))
	app.Post("/:userId", updateUser(service))
	app.Delete("/:userId", deleteUser(service))
	app.Post("/:userId/pets", ownPets(service))
}

func createUser(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		var user *entity.User
		err := c.BodyParser(&user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user, err = service.CreateUser(ctx, user.Email, user.Password, user.FirstName, user.LastName)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func getUser(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := service.GetUser(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User Found",
			"data":    toJ,
		})
	}
}

func updateUser(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var user *entity.User

		err = c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user.ID = id
		user, err = service.UpdateUser(ctx, user)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func deleteUser(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "Bad Id Format",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		err = service.DeleteUser(ctx, &id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error deleting user",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "user deleted successfully",
			"error":  nil,
		})
	}
}

func listUsers(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		users, err := service.ListUsers(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := make([]presenter.User, len(users))

		for i, user := range users {
			toJ[i] = presenter.User{
				ID:        user.ID,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    toJ,
		})
	}
}

func ownPets(service user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		id, err := entity.StringToID(c.Params("userId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var petIDs []*entity.ID

		err = c.BodyParser(&petIDs)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		err = service.OwnPets(ctx, &id, petIDs)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
		})
	}
}
