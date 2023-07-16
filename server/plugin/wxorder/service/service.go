package service
import (
   "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxorder/model"
)


type WxorderService struct{}

func (e *WxorderService) PlugService(req model.Request ) (res model.Response,err error) {
    // 写你的业务逻辑
	return res, nil
}
