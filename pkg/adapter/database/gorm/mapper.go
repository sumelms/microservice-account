package gorm

import (
	domain "github.com/sumelms/microservice-account/pkg/domain"
)

func toDBModel(entity *domain.User) *User {
	user := &User{
		ID:        entity.ID,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if entity.DeletedAt != nil {
		user.DeletedAt = entity.DeletedAt
	}
	if entity.ActivatedAt != nil {
		user.ActivatedAt = entity.ActivatedAt
	}

	return user
}

func toDomainModel(entity *User) *domain.User {
	return &domain.User{
		ID:          entity.ID,
		Email:       entity.Email,
		Password:    entity.Password,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		ActivatedAt: entity.ActivatedAt,
	}
}
