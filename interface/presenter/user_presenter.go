package presenter

import "tutorial/domain/model"

// UserInterfacePresenter is an interface for user presenter
// Path: interface\repository\user_presenter.go
// Used to manage response data from usecase to interactor

type UserInterfacePresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
