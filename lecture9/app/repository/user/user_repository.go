package user

import (
	"errors"
	"justCode/lecture9/app/model"
)

type UserRepository interface {
	GetAll() []model.User
	GetByID(id int) (*model.User, error)
	Create(user *model.User) error
	Update(id int, user *model.User) error
	Delete(id int) error
}

// IMPORTANT: Эту сущность я использую чтобы хранить данные
type UserRepositoryStr struct {
	users []model.User
}

func NewUserRepository() *UserRepositoryStr {
	return &UserRepositoryStr{}
}

func (r *UserRepositoryStr) GetAll() []model.User {
	return r.users
}

func (r *UserRepositoryStr) GetByID(id int) (*model.User, error) {
	for i := range r.users {
		if r.users[i].Id == id {
			return &r.users[i], nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *UserRepositoryStr) Create(user *model.User) error {
	user.Id = len(r.users) + 1
	r.users = append(r.users, *user)
	return nil
}

func (r *UserRepositoryStr) Update(id int, user *model.User) error {
	for i := range r.users {
		if r.users[i].Id == id {
			r.users[i] = *user
			return nil
		}
	}
	return errors.New("Not found")
}

func (r *UserRepositoryStr) Delete(id int) error {
	for i := range r.users {
		if r.users[i].Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}
