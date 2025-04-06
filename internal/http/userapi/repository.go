package userapi

import (
	"context"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}

type InMemoryUserRepository struct {
	users map[int64]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int64]*User),
	}
}

func (r *InMemoryUserRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	var users []*User

	//check the map if it is ok
	//fmt.Println("the message from repo/getallusers is:", len(r.users), r.users[0].Name)

	for id := 0; id < len(r.users); id++ {
		users = append(users, r.users[int64(id)])
	}

	return users, nil
}

func (r *InMemoryUserRepository) GetUserById(ctx context.Context, id int64) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user *User) error {
	r.users[user.ID] = user
	fmt.Println("the message from /repo/createuser is:", user)
	return nil
}
