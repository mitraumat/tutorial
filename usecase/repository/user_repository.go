package repository

import (
	"gorm.io/gorm"
	"tutorial/domain/model"
	"tutorial/domain/request"
	"tutorial/interface/repository"
)

type userRepository struct {
	db *gorm.DB
}

func (ur userRepository) DeleteUser(u *model.User) (*model.User, error) {
    err := ur.db.Delete(&u).Error
    if err != nil {
        return nil, err
    }

    return u,nil
}

func (ur userRepository) UpdateUser(u *model.User, r *request.UserUpdateRequest) (*model.User, error) {
	u.Username = r.Username
	u.Name = r.Name
	err := ur.db.Save(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur userRepository) FindById(u *model.User, query string) (*model.User, error) {
	err := ur.db.First(&u, "id = ?", query).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur userRepository) CreateUser(r *model.User) (*model.User, error) {
	err := ur.db.Create(&r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (ur userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UseCaseUserPresenter(db *gorm.DB) repository.UserInterfaceRepository {
	return &userRepository{db}
}
