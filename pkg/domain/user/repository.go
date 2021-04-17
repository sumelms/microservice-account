package user

type RepositoryInterface interface {
	Store(user *User) (*User, error)
	GetByID(id string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	GetAll() ([]User, error)
}
