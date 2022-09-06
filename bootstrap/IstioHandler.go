package bootstrap

import (
	"istiomang/pkg/vs"
)

//注入 回调handler
type IstioHandler struct{}

func NewIstioHandler() *IstioHandler {
	return &IstioHandler{}
}

// VsHandler handler
func (this *IstioHandler) VsHandler() *vs.VsHandler {
	return &vs.VsHandler{}
}
