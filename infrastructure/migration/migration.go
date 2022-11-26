package migration

import (
    "tutorial/domain/model"
    "tutorial/infrastructure/datastore"
)

func Run() {
//	err := datastore.Database().AutoMigrate(&user.User{})
    err := datastore.Database().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
