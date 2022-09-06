package v1

import "fmt"

type IUserRepo interface {
	FindById(model *UserModel) error
	New(model *UserModel) error
	Update(model *UserModel) error
}

type UserRepo struct {
}

func (u *UserRepo) FindById(model *UserModel) error {
	if model.UserName == "shenyi" {
		model.UserPass = "123"
	} else if model.UserName == "list" {
		model.UserPass = "234"
	} else {
		return fmt.Errorf("no soch user %s", model.UserName)
	}
	return nil
}

func (u *UserRepo) New(model *UserModel) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Update(model *UserModel) error {
	//TODO implement me
	panic("implement me")
}
