package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

var users = []User{
	{ID: uuid.NewString(), Name: "John Doe", CreatedAt: time.Now().UTC().Format(time.RFC3339)},
	{ID: uuid.NewString(), Name: "Jane Doe", CreatedAt: time.Now().UTC().Format(time.RFC3339)},
	{ID: uuid.NewString(), Name: "John Smith", CreatedAt: time.Now().UTC().Format(time.RFC3339)},
}

type Service interface {
	GetAll() []User
	GetByID(id string) (User, error)
	Create(user User) (User, error)
	Update(id string, user User) (User, error)
	Delete(id string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GetAll() []User {
	return users
}

func (s *service) GetByID(id string) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}

func (s *service) Create(user User) (User, error) {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	users = append(users, user)
	return user, nil
}

func (s *service) Update(id string, user User) (User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			return user, nil
		}
	}

	return User{}, nil
}

func (s *service) Delete(id string) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return nil
}
