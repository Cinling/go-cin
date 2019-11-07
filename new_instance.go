package cin

import (
	"github.com/cinling/cin/components"
	"github.com/cinling/cin/configs"
	"github.com/cinling/cin/utils"
)

// 新建 websocket server 配置
func NewConfigWebsocketServer(port uint16) *configs.WebsocketServer {
	config := new(configs.WebsocketServer)
	config.Port = port
	config.Component = &components.WebsocketServer{}
	config.MessageParseHandler = nil
	return config
}

// 新建 http server 配置
func NewConfigHttpServer(port uint16) *configs.HttpServer {
	config := new(configs.HttpServer)
	config.Component = &components.HttpServer{}
	config.Port = port
	config.SessionKey = "cin-key"
	config.SessionName = "cin"
	return config
}

// 新建 数据库配置
func NewConfigDatabase(driverName string, host string, port string, name string, username string, password string) *configs.Database {
	config := new(configs.Database)
	config.Component = &components.Database{}
	config.DriverName = driverName
	config.Host = host
	config.Port = port
	config.Name = name
	config.Username = username
	config.Password = password
	config.SetMigrateDir(utils.File.GetRunPath() + "/database/migrations")
	return config
}

// 新建 控制台配置
func NewConfigConsole() *configs.Console {
	config := new(configs.Console)
	config.Component = &components.Console{}
	return config
}
