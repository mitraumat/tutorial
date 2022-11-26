package registry

import (
	"gorm.io/gorm"
	"tutorial/infrastructure/helper"
)

type registry struct {
	db *gorm.DB
}

func (r *registry) RegistryAppController() helper.Application {
	return helper.Application{
		User: r.RegistryUserController(),
	}
}

type Registry interface {
	RegistryAppController() helper.Application
}

func BaseRegistry(db *gorm.DB) Registry {
	return &registry{db}

}
