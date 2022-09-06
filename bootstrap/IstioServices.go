package bootstrap

import "istiomang/pkg/vs"

//@Config
type IstioServiceConfig struct{}

func NewIstioServiceConfig() *IstioServiceConfig {
	return &IstioServiceConfig{}
}
func (*IstioServiceConfig) VsService() *vs.VsService {
	return vs.NewVsService()
}
