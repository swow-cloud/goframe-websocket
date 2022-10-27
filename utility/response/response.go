package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type JsonRes struct {
	Code    int         `json:"code"`    //错误码(0:成功;1:失败)
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    //返回数据
}

// Json 返回标准json数据
func Json(request *ghttp.Request, code int, message string, data interface{}) {

	var responseData interface{}
	if len(message) > 0 {
		responseData = data
	} else {
		responseData = g.Map{}
	}
	request.Response.WriteJson(JsonRes{Code: code,
		Message: message,
		Data:    responseData})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(request *ghttp.Request, code int, message string, data interface{}) {
	Json(request, code, message, data)
	request.Exit()
}
