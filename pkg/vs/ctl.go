package vs

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type VsCtl struct {
	VsService *VsService `inject:"-"`
}

func NewVsCtl() *VsCtl {
	return &VsCtl{}
}
func (this *VsCtl) VsList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": this.VsService.ListVs(ns),
	}
}
func (*VsCtl) Name() string {
	return "VsCtl"
}
func (this *VsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/virtualservices", this.VsList)
}
