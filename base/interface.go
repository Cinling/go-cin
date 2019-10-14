package base

import "net/http"

// 框架内接口

// 组件类型配置接口
type ConfigComponentInterface interface {
	GetComponent() ComponentInterface
}

// 组件接口
type ComponentInterface interface {
	// 初始化
	Init(configInterface ConfigComponentInterface)
	// 开始
	Start()
	// 停止
	Stop()
}

// 控制器实例
type ControllerInterface interface {
	// 初始化方法
	Init()
	// 执行动作前执行的方法
	BeforeAction(action string) bool
	// 执行动作后执行的方法
	AfterAction(action string, response []byte) []byte

	// 设置上下文对象
	SetContext(context ContextInterface)
	// 获取上下文对象
	GetContext() ContextInterface

	// 设置 http 请求的参数
	SetHttpValues(w http.ResponseWriter, r *http.Request)
	// 设置 websocket 请求的参数
	//SetWebsocketValues()
}

// 上下文接口
type ContextInterface interface {
	// 获取 session
	GetSession() SessionInterface
}

// session 接口
type SessionInterface interface {
	// 获取 sessionId
	GetSessionId() string
	// 设置值
	Set(key interface{}, value interface{})
	// 获取值
	Get(key interface{}) interface{}
	// 销毁session
	Destroy()
}