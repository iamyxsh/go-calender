package database

import (
	"calendly/constants/configs"
	errorhandler "calendly/utils/error"
	"fmt"

	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	err := mgm.SetDefaultConfig(nil, "Calender", options.Client().ApplyURI(configs.MONGO_URI))

	errorhandler.HandleErr(err)

	fmt.Print("Database Connected....")
}
