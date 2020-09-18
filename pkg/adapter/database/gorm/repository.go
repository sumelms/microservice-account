package gorm

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	domain "github.com/sumelms/microservice-user/pkg/domain"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) *Repository {
	db.AutoMigrate(&User{})

	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r Repository) Store(user *domain.User) (*domain.User, error) {
	entity := toDBModel(user)

	if err := r.db.Create(entity).Error; err != nil {
		return nil, errors.Wrap(err, "Repository.CreateUser")
	}

	return toDomainModel(entity), nil
}

func (r Repository) GetByID(id string) (*domain.User, error) {
	var result User

	query := r.db.Where("id = ?", id).First(&result)

	if query.RecordNotFound() {
		return nil, errors.New("User not found")
	}

	if err := query.Error; err != nil {
		return nil, errors.Wrap(err, "Repository.GetUserByID")
	}

	return toDomainModel(&result), nil
}

func (r Repository) Update(entity *domain.User) (*domain.User, error) {
	// FIXME Can we improve the update process?
	var user User

	id := entity.ID
	query := r.db.Where("id = ?", id).First(&user)

	if query.RecordNotFound() {
		return nil, errors.New("User not found")
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

func (r Repository) Delete(id string) error {
	query := r.db.Model(User{}).Where("id = ?", id).Update("deleted_at", "NOW()")

	if err := query.Error; err != nil {
		return err
	}

	return nil
}

func (r Repository) GetAll() ([]domain.User, error) {
	var results []User

	if err := r.db.Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "Repository.ListUsers")
	}

	users := make([]domain.User, len(results))

	for index, element := range results {
		users[index] = *toDomainModel(&element)
	}

	return users, nil
}
