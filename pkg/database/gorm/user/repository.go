package user

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	domain "github.com/sumelms/microservice-account/pkg/domain/user"
)

const whereQuery = "id = ?"

// Repository struct
type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&User{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

// Store creates an user
func (r Repository) Store(user *domain.User) (*domain.User, error) {
	entity := toDBModel(user)

	if err := r.db.Create(entity).Error; err != nil {
		return nil, errors.Wrap(err, "Repository.CreateUser")
	}

	return toDomainModel(entity), nil
}

// GetByID get a user by its ID
func (r Repository) GetByID(id string) (*domain.User, error) {
	var result User

	query := r.db.Where(whereQuery, id).First(&result)

	if query.RecordNotFound() {
		return nil, errors.New("user not found")
	}

	if err := query.Error; err != nil {
		return nil, errors.Wrap(err, "Repository.GetUserByID")
	}

	return toDomainModel(&result), nil
}

// Update the given user
func (r Repository) Update(entity *domain.User) (*domain.User, error) {
	var user User

	id := entity.ID.String()
	query := r.db.Where(whereQuery, id).First(&user)

	if query.RecordNotFound() {
		return nil, errors.New("user not found")
	}

	if entity.Email != "" {
		user.Email = entity.Email
	}

	if entity.Password != "" {
		user.Password = entity.Password
	}

	query = r.db.Save(&user)

	if err := query.Error; err != nil {
		return nil, errors.Wrap(err, "Repository.Update")
	}

	return toDomainModel(&user), nil
}

// Delete a user by ID
func (r Repository) Delete(id string) error {
	query := r.db.Model(User{}).Where(whereQuery, id).Update("deleted_at", "NOW()")

	if err := query.Error; err != nil {
		return err
	}

	return nil
}

// GetAll returns a list of users
func (r Repository) GetAll() ([]domain.User, error) {
	var results []User

	if err := r.db.Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "Repository.ListUsers")
	}

	users := make([]domain.User, len(results))

	for index := range results {
		users[index] = *toDomainModel(&results[index])
	}

	return users, nil
}
