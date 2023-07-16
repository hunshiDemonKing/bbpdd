package wxorder

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxorder/router"
	"github.com/gin-gonic/gin"
)

type WxorderPlugin struct {
}

func CreateWxorderPlug()*WxorderPlugin {
	return &WxorderPlugin{}
}

func (*WxorderPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitWxorderRouter(group)
}

func (*WxorderPlugin) RouterPath() string {
	return "wxorder"
}
