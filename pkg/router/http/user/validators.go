package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	domain2 "github.com/sumelms/sumelms/user/pkg/domain"
)

type CreateUserValidator struct {
	Username        string `http:"username" validate:"required"`
	Email           string `http:"email" validate:"required,email"`
	Password        string `http:"password" validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `http:"confirm_password" validate:"required`
}

func BindCreateUser(c echo.Context) (*domain2.User, error) {
	entity := new(CreateUserValidator)

	if err := c.Bind(entity); err != nil {
		return nil, err
	}

	if err := c.Validate(entity); err != nil {
		return nil, err
	}

	user := &domain2.User{
		Username: entity.Username,
		Email:    entity.Email,
		Password: entity.Password,
	}

	return user, nil
}

type UpdateUserValidator struct {
	Email           string `http:"email", validate:"required,email"`
	Password        string `http:"password" validate:"eqfield=ConfirmPassword"`
	ConfirmPassword string `http:"confirm_password" validate="required_with=Password"`
}

func BindUpdateUser(c echo.Context) (*domain2.User, error) {
	entity := new(UpdateUserValidator)

	if err := c.Bind(entity); err != nil {
		return nil, err
	}

	if err := c.Validate(entity); err != nil {
		return nil, err
	}

	// The URL param
	id := c.Param("id")

	user := &domain2.User{
		ID:    uuid.MustParse(id),
		Email: entity.Email,
	}

	if entity.Password != "" {
		user.Password = entity.Password
	}

	return user, nil
}
