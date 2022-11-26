package registry

import (
	ic "tutorial/interface/controller"
	ii "tutorial/interface/interactor"
	ip "tutorial/interface/presenter"
	ir "tutorial/interface/repository"
	uc "tutorial/usecase/controller"
	ui "tutorial/usecase/interactor"
	up "tutorial/usecase/presenter"
	ur "tutorial/usecase/repository"
)

func (r *registry) RegistryUserController() ic.UserInterfaceController {
	return uc.UseCaseUserController(r.RegistryUserInteractor())
}

func (r *registry) RegistryUserInteractor() ii.UserInterfaceInteractor {
	return ui.UseCaseUserInteractor(r.RegistryUserRepository(), r.RegistryUserPresenter())
}

func (r *registry) RegistryUserRepository() ir.UserInterfaceRepository {
	return ur.UseCaseUserPresenter(r.db)
}

func (r *registry) RegistryUserPresenter() ip.UserInterfacePresenter {
	return up.UseCaseUserPresenter()
}
