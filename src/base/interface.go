package base
// 框架内接口

// 组件类型配置接口
type ConfigComponentInterface interface {
	GetComponent() ComponentInterface
}

// 组件接口
type ComponentInterface interface {
	Init(configInterface ConfigComponentInterface)
	Start()
	Run()
	Stop()
	Destroy()
}