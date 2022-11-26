package presenter

import (
	"tutorial/domain/model"
	"tutorial/interface/presenter"
)

type userPresenter struct{}

func (up userPresenter) ResponseUsers(u []*model.User) []*model.User {
	return u
}

func UseCaseUserPresenter() presenter.UserInterfacePresenter {
	return &userPresenter{}
}
