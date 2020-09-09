package user

import (
	"github.com/pkg/errors"
)

type Service interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
	ListUsers() ([]User, error)
}

type userService struct {
	repo Repository
}

func NewService(repository Repository) *userService {
	return &userService{
		repo: repository,
	}
}

func (s userService) CreateUser(user *User) (*User, error) {
	user, err := s.repo.Store(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.CreateUser")
	}
	return user, nil
}

func (s userService) GetUserByID(id string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "Service.GetUserByID")
	}
	return user, nil
}

func (s userService) UpdateUser(user *User) (*User, error) {
	updated, err := s.repo.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.UpdateUser")
	}
	return updated, nil
}

func (s userService) DeleteUser(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "Service.DeleteUser")
	}
	return nil
}

func (s userService) ListUsers() ([]User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "Serive.ListUsers")
	}
	return users, nil
}
