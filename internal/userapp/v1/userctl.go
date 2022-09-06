package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"net/http"
)

type UserCtl struct {
	UserService *UserService `inject:"-"`
}

func UserController() *UserCtl {
	return &UserCtl{}
}

func (u *UserCtl) Login(c *gin.Context) goft.Json {
	req := &KindLoginRequest{}
	goft.Error(c.ShouldBindJSON(req))
	rsp, err := u.UserService.UserLogin(c, req)
	goft.Error(err)
	return rsp
}

func (u *UserCtl) Name() string {
	return "UserCtl"
}

func (u *UserCtl) Build(goft *goft.Goft) {
	goft.Handle(http.MethodPost, "/login", u.Login)
}
