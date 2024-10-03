package bootstrap

import (
	"github.com/sabareeswaran-sakthivel/e-commerce/database/connection"
	"github.com/sabareeswaran-sakthivel/e-commerce/routes"
	commonUtils "github.com/sabareeswaran-sakthivel/e-commerce/utils"
)

func Init() {

	connection.InitSqliteConnection()

	commonUtils.CreateTables()

	go commonUtils.RunCronJob()

	commonUtils.InsertRegions()

	routes.Init()

}
