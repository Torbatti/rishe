package main

import (
	"github.com/torbatti/rishe"
	"github.com/torbatti/rishe/cogs"
	"github.com/torbatti/rishe/modules"
)

func main() {
	const app_name = ""
	const app_version = ""

	app := rishe.New()

	// Modules
	modules.Broker_Default()
	modules.Cache_Redis()
	modules.Cmd_Default()
	modules.Cron_Default()
	modules.Db_Sqlite()
	modules.Store_Default()
	modules.View_Default()

	// Default Middlewares
	cogs.Middleware_logger()
	cogs.Middleware_Recover()
	cogs.Middleware_Cors()
	cogs.Middleware_Compress()
	cogs.Middleware_HttpRateLimit()

	// Auth Cogs
	cogs.Auth_Migration()
	cogs.Authenicate_Middleware()
	cogs.Auth_Route()
	cogs.Auth_Api_Json()
	cogs.Auth_View_Login()
	cogs.Auth_View_Register()

	app.Start()
}
