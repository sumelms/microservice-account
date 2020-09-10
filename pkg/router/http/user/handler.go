package user

import (
	"github.com/labstack/echo"
	"github.com/sumelms/sumelms/user/pkg/context"
	"net/http"
)

func NewHandler(e *echo.Echo) {
	e.GET("/users", listUsers)
	e.GET("/users/:id", getUserByID)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

func createUser(c echo.Context) error {
	entity, err := BindCreateUser(c)
	if err != nil {
		return err
	}

	user, err := c.(*context.Context).Service.CreateUser(entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := c.(*context.Context).Service.GetUserByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func listUsers(c echo.Context) error {
	users, err := c.(*context.Context).Service.ListUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	err := c.(*context.Context).Service.DeleteUser(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func updateUser(c echo.Context) error {
	entity, err := BindUpdateUser(c)
	if err != nil {
		return err
	}

	updatedUser, err := c.(*context.Context).Service.UpdateUser(entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedUser)
}
