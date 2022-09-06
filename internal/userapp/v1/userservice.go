package v1

import (
	"context"
	"fmt"
)

type UserService struct {
	UnimplementedUserServiceServer

	UserRepo IUserRepo `inject:"-"`
}

func (u *UserService) UserLogin(ctx context.Context, request *KindLoginRequest) (*KindLoginResponse, error) {
	user := &UserModel{
		UserName: request.Spec.UserName,
	}
	err := u.UserRepo.FindById(user)
	if err != nil {
		return nil, err
	}

	if user.UserPass != request.Spec.UserPass {
		return nil, fmt.Errorf("error username or password")
	}

	return &KindLoginResponse{
		Code:    200,
		Message: "success",
		Data: &UserModel{
			UserId:   101,
			UserName: request.Spec.UserName,
		},
	}, nil
}
