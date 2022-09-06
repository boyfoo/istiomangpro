package bootstrap

import "istiomang/pkg/vs"

type IstioMaps struct {
}

func NewIstioMaps() *IstioMaps {
	return &IstioMaps{}
}

//初始化 VsMapStruct
func (this *IstioMaps) InitVsMap() *vs.VsMapStruct {
	return &vs.VsMapStruct{}
}
