package v1

type UserV1Config struct {
}

func NewUserV1Config() *UserV1Config {
	return &UserV1Config{}
}

func (u *UserV1Config) InitUserService() *UserService {
	return &UserService{}
}

func (u *UserV1Config) InitUserRepo() *UserRepo {
	return &UserRepo{}
}
