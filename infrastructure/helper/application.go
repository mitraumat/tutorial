package helper

import "tutorial/interface/controller"

type Application struct {
	User interface {
		controller.UserInterfaceController
	}
}
