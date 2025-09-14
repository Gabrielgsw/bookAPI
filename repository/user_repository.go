package repository

type UserRepository interface {
	GetAll() ([]model.User, error)
}
