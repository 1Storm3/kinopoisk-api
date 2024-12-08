package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"kbox-api/internal/modules/auth/dto"
	"kbox-api/internal/modules/auth/mapper"
	"kbox-api/internal/modules/auth/service"
	"kbox-api/shared/httperror"
)

type AuthHandlerInterface interface {
	Login(ctx *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Me(c *fiber.Ctx) error
}

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{
		authService: authService,
	}
}

// Login @Summary Login user
// @Description Login to the application
// @Accept  json
// @Produce  json
// @Param login body dto.LoginDTO true "Login information"
// @Success 200 token string
// @Failure 400 {object} httperror.Error
// @Router /auth/login [post]
func (a *AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginData dto.LoginDTO
	if err := ctx.BodyParser(&loginData); err != nil {

		return httperror.New(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	token, err := a.authService.Login(loginData)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {
	var requestUser dto.RegisterDTO
	if err := c.BodyParser(&requestUser); err != nil {
		return httperror.New(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	tokenString, err := a.authService.Register(requestUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
func (a *AuthHandler) Me(c *fiber.Ctx) error {
	result, err := a.authService.Me(c)
	if err != nil {
		return err
	}

	user := mapper.MapModelUserToResponseDTO(result)

	return c.JSON(user)
}
