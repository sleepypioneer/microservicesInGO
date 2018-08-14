package main

import (
	"fmt"

	"github.com/sleepypioneer/microservicesInGO/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	services.StartWebServer("6767")
}
