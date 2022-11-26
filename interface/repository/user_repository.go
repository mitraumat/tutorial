package repository

import (
	"tutorial/domain/model"
    "tutorial/domain/request"
)

// UserRepository is a repository interface for user
// Path: interface\repository\user_repository.go
// Used to manage the data layer of the application from database

type UserInterfaceRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	CreateUser(r *model.User) (*model.User, error)
	FindById(u *model.User, query string) (*model.User, error)
    UpdateUser(u *model.User,r *request.UserUpdateRequest)(*model.User,error)
    DeleteUser(u *model.User) (*model.User,error)
}
