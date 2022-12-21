package app

import (
	"Diploma/app/modules"
	"Diploma/app/server"
	"Diploma/pkg/config"
	"Diploma/pkg/database"
	"github.com/jmoiron/sqlx"
	"log"
)

func Run() {
	defer log.Print("Shutting down\n")

	//Logging initialisation
	//if err := logging.LogInit(); err != nil {
	//	log.Fatal("Logging init error\n" + err.Error())
	//}

	//Config initialisation
	if err := config.Init(); err != nil {
		log.Fatal("Config init error\n" + err.Error())
	}

	//Connecting to Postgres
	PostgresClient, err := database.Setup()
	if err != nil {
		log.Fatal("Db connect error\n", err.Error())
	}

	//SetupModules
	handlers := SetupModules(PostgresClient)

	//Start handling http requests
	if err := server.Start(handlers); err != nil {
		log.Fatal("Server starting error\n" + err.Error())
	}

}

func SetupModules(PostgresClient *sqlx.DB) (handlers modules.HandlerModule) {
	GatewayModule := modules.SetupGateway(PostgresClient)
	log.Println("Gateway module setup correctly")
	DriveModule := modules.SetupDrive()
	log.Println("Drive module setup correctly")
	UseCaseModule := modules.SetupUseCase(GatewayModule, DriveModule)
	log.Println("UseCase module setup correctly")
	DelegateModule := modules.SetupDelegate(UseCaseModule)
	log.Println("Delegate module setup correctly")
	HandlerModule := modules.SetupHandler(DelegateModule)
	log.Println("Handler module setup correctly")
	return HandlerModule
}
