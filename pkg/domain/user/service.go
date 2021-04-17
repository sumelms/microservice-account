package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type ServiceInterface interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]User, error)
}

type Service struct {
	repo   RepositoryInterface
	logger log.Logger
}

func NewService(repository RepositoryInterface, logger log.Logger) *Service {
	return &Service{
		repo:   repository,
		logger: logger,
	}
}

func (s Service) CreateUser(_ context.Context, user *User) (*User, error) {
	user, err := s.repo.Store(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.CreateUser")
	}
	return user, nil
}

func (s Service) GetUser(_ context.Context, id string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "Service.GetUserByID")
	}
	return user, nil
}

func (s Service) UpdateUser(_ context.Context, user *User) (*User, error) {
	updated, err := s.repo.Update(user)
	if err != nil {
		return nil, errors.Wrap(err, "Service.UpdateUser")
	}
	return updated, nil
}

func (s Service) DeleteUser(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "Service.DeleteUser")
	}
	return nil
}

func (s Service) ListUsers(_ context.Context) ([]User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "Serive.ListUsers")
	}
	return users, nil
}
