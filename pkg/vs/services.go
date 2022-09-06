package vs

import "istio.io/client-go/pkg/apis/networking/v1alpha3"

//@Service
type VsService struct {
	VsMap *VsMapStruct `inject:"-"`
}

func NewVsService() *VsService {
	return &VsService{}
}
func (this *VsService) ListVs(ns string) []*v1alpha3.VirtualService {
	return this.VsMap.ListAll(ns)
}
