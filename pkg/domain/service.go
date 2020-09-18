package domain

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type Service interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]User, error)
}

type userService struct {
	repo   Repository
	logger log.Logger
}

func NewService(repository Repository, logger log.Logger) *userService {
	return &userService{
		repo:   repository,
		logger: logger,
	}
}

func (s userService) CreateUser(_ context.Context, user *User) (*User, error) {
	user, err := s.repo.Store(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.CreateUser")
	}
	return user, nil
}

func (s userService) GetUser(_ context.Context, id string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "Service.GetUserByID")
	}
	return user, nil
}

func (s userService) UpdateUser(_ context.Context, user *User) (*User, error) {
	updated, err := s.repo.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.UpdateUser")
	}
	return updated, nil
}

func (s userService) DeleteUser(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "Service.DeleteUser")
	}
	return nil
}

func (s userService) ListUsers(_ context.Context) ([]User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "Serive.ListUsers")
	}
	return users, nil
}
