package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	v1 "istiomang/internal/userapp/v1"
)

func main() {
	//svc := &v1.UserService{}
	//fmt.Println(svc.UserLogin(context.Background(), &v1.KindLoginRequest{
	//	Spec: &v1.UserLoginModel{
	//		UserName: "zx",
	//		UserPass: "123",
	//	},
	//}))

	goft.Ignite().Config(v1.NewUserV1Config()).Mount("", v1.UserController()).Launch()

}
