package handlers

import (
	"context"
	"fmt"
	"github.com/aurareaper/event-management-app/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"time"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	token, user, err := h.service.Login(context, creds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully logged in",
		"data": &fiber.Map{
			"token":          token,
			"user":           user,
			"availableRoles": models.RoleOptions(),
		},
	})
}

func (h *AuthHandler) Register(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": fmt.Errorf("please provide a valid email and password").Error(),
		})
	}

	// Validate role if provided
	if creds.Role != "" {
		if creds.Role != models.Manager && creds.Role != models.Attendee {
			roleOptions := models.RoleOptions()
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":     "fail",
				"message":    fmt.Sprintf("Invalid role: '%s'. Must be one of: %v", creds.Role, roleOptions),
				"validRoles": roleOptions,
			})
		}
	}

	token, user, err := h.service.Register(context, creds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "Successfully registered",
		"data": &fiber.Map{
			"token":          token,
			"user":           user,
			"availableRoles": models.RoleOptions(),
		},
	})
}

func NewAuthHandler(route fiber.Router, service models.AuthService) {
	handler := &AuthHandler{
		service: service,
	}

	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)
}
