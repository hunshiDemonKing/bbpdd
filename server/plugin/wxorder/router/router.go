package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxorder/api"
	"github.com/gin-gonic/gin"
)

type WxorderRouter struct {
}

func (s *WxorderRouter) InitWxorderRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WxorderApi
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
