package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID
	Username    string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ActivatedAt *time.Time
}
