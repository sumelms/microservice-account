package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	domain "github.com/sumelms/sumelms/user/pkg/domain/user"
)

type CreateUserValidator struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required`
}

func BindCreateUser(c echo.Context) (*domain.User, error) {
	entity := new(CreateUserValidator)

	if err := c.Bind(entity); err != nil {
		return nil, err
	}

	if err := c.Validate(entity); err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: entity.Username,
		Email:    entity.Email,
		Password: entity.Password,
	}

	return user, nil
}

type UpdateUserValidator struct {
	Email           string `json:"email", validate:"required,email"`
	Password        string `json:"password" validate:"eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" validate="required_with=Password"`
}

func BindUpdateUser(c echo.Context) (*domain.User, error) {
	entity := new(UpdateUserValidator)

	if err := c.Bind(entity); err != nil {
		return nil, err
	}

	if err := c.Validate(entity); err != nil {
		return nil, err
	}

	// The URL param
	id := c.Param("id")

	user := &domain.User{
		ID:    uuid.MustParse(id),
		Email: entity.Email,
	}

	if entity.Password != "" {
		user.Password = entity.Password
	}

	return user, nil
}
