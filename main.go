package main

import (
	"fmt"

	"github.com/sleepypioneer/microservicesInGO/dbclient"
	"github.com/sleepypioneer/microservicesInGO/services"
)

var appName = "MicroservicesInGO"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	services.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	services.DBClient = dbclient.BoltClient{}
	services.DBClient.OpenBoltDb()
	services.DBClient.Seed()
	// Our microservice should now create a database on start.
}
