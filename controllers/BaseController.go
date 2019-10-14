package controllers

import (
	"cin/base"
	"net/http"
)

// 所有 handler 的基类。主要处理问题是：统一接口、数据库管理
type BaseController struct {
	base.ControllerInterface

	// 上下文
	context base.ContextInterface
	// 接收到请求的参数
	values map[string]interface{}
}

// 初始化
func (controller *BaseController) Init() {
	controller.values = map[string]interface{}{}
}

// 执行动作前执行的方法
// 如果返回 false 将会返回一个错误
func (controller *BaseController) BeforeAction(action string) bool {
	return true
}

// 执行动作后执行的方法
// 过滤返回结果
func (controller *BaseController) AfterAction(action string, response []byte) []byte {
	return response
}

// 设置上下文对象
func (controller *BaseController) SetContext(context base.ContextInterface) {
	controller.context = context
}
// 获取上下文对象
func (controller *BaseController) GetContext() base.ContextInterface {
	return controller.context
}

// 设置http接收到的参数
func (controller *BaseController) SetHttpValues(w http.ResponseWriter, r *http.Request) {
	// 接收 get 和 post 参数
	_ = r.ParseForm()
	for key, value := range r.Form {
		controller.values[key] = value
	}

	// TODO 处理数组、对象、数组和对象混合的数据
}