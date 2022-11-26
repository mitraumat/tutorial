package interactor

import (
	"tutorial/domain/model"
    "tutorial/domain/request"
)

// UserInterfaceInteractor is an interface for user interactor
// Path: interface\interactor\user_interactor.go
// Used to manage usecase from interactor to controller

type UserInterfaceInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(r *model.User) (*model.User, error)
	GetById(u *model.User, id string) (*model.User, error)
    Update(u *model.User,r * request.UserUpdateRequest)(*model.User,error)
    DeleteById(u *model.User) (*model.User,error)
}
