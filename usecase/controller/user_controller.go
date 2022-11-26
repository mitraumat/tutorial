package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"tutorial/domain/model"
    "tutorial/domain/request"
    "tutorial/infrastructure/helper"
	"tutorial/interface/controller"
	"tutorial/interface/interactor"
)

type userController struct {
	userInteractor interactor.UserInterfaceInteractor
}

func (uc userController) DeleteById(context *fiber.Ctx) error {
    var u *model.User
    id := context.Params("id")

    find,err := uc.userInteractor.GetById(u,id)

    if err != nil {
        log.Println(err)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, "DATA_NOT_FOUND"))
    }

    deleted,err := uc.userInteractor.DeleteById(find)

    if err != nil {
        log.Println(err)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, "FAILED"))
    }
    log.Println(deleted)
    return helper.ResponseJson(context,helper.NewResponse(fiber.StatusOK,true,"HTTP_OKE"))

}

func (uc userController) UpdateUser(context *fiber.Ctx) error {
    var u *model.User
    req := new(request.UserUpdateRequest)

    id := context.Params("id")
    find, errs := uc.userInteractor.GetById(u, id)

    if errs != nil {
        log.Println(errs)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, "DATA_NOT_FOUND"))
    }


    if err := context.BodyParser(req); err != nil {
        log.Println(err)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, err.Error()))
    }

    errors := helper.ValidateStruct(req)
    if errors != nil {
        log.Print(errors)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, "Validation Error", helper.WithData(errors)))
    }

    u,err := uc.userInteractor.Update(find,req)

    if err != nil {
        log.Println(err)
        return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, err.Error()))
    }

    return helper.ResponseJson(context,helper.NewResponse(fiber.StatusOK,true,"HTTP_OKE",helper.WithData(find)))
}

func (uc userController) GetUserById(context *fiber.Ctx) error {
	var u *model.User
	id := context.Params("id")
	u, err := uc.userInteractor.GetById(u, id)
	if err != nil {
		log.Println(err)
		return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, err.Error()))
	}
    return helper.ResponseJson(context, helper.NewResponse(fiber.StatusOK, true, "HTTP_OK", helper.WithData(u)))

}

func (uc userController) GetUsers(context *fiber.Ctx) error {
	var u []*model.User
	u, err := uc.userInteractor.Get(u)
	if err != nil {
		log.Println(err)
		return helper.ResponseJson(context, helper.NewResponse(http.StatusBadRequest, false, "Internal Server Error"))
	}
	return helper.ResponseJson(context, helper.NewResponse(http.StatusOK, true, "HTTP_OK", helper.WithData(u)))
}

func (uc userController) CreateUser(context *fiber.Ctx) error {
	req := new(model.User)

	if err := context.BodyParser(req); err != nil {
		log.Println(err)
		return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, err.Error()))
	}
	errors := helper.ValidateStruct(req)
	if errors != nil {
		return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, "Validation Error", helper.WithData(errors)))
	}
	u, err := uc.userInteractor.Create(req)
	if err != nil {
		log.Println(err)
		return helper.ResponseJson(context, helper.NewResponse(fiber.StatusBadRequest, false, err.Error()))
	}
	log.Println(u)
	return helper.ResponseJson(context, helper.NewResponse(fiber.StatusOK, true, "HTTP_OK", helper.WithData(u)))
}

func UseCaseUserController(us interactor.UserInterfaceInteractor) controller.UserInterfaceController {
	return &userController{us}
}
