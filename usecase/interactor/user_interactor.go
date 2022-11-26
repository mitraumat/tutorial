package interactor

import (
	"tutorial/domain/model"
    "tutorial/domain/request"
    "tutorial/interface/interactor"
	"tutorial/interface/presenter"
	"tutorial/interface/repository"
)

type userInteractor struct {
	UserRepository repository.UserInterfaceRepository
	UserPresenter  presenter.UserInterfacePresenter
}

func (us userInteractor) DeleteById(u *model.User) (*model.User, error) {
    user , err := us.UserRepository.DeleteUser(u)
    if err != nil {
        return nil,err
    }
    return user,nil
}

func (us userInteractor) Update(u *model.User, r *request.UserUpdateRequest) (*model.User, error) {
    user, err := us.UserRepository.UpdateUser(u,r)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (us userInteractor) GetById(u *model.User, id string) (*model.User, error) {
	user, err := us.UserRepository.FindById(u, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us userInteractor) Create(r *model.User) (*model.User, error) {
	user, err := us.UserRepository.CreateUser(r)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (us userInteractor) Get(u []*model.User) ([]*model.User, error) {
	users, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUsers(users), nil
}

func UseCaseUserInteractor(r repository.UserInterfaceRepository, p presenter.UserInterfacePresenter) interactor.UserInterfaceInteractor {
	return &userInteractor{r, p}
}
