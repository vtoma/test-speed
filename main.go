package main

import (
	speedTestLib "github.com/vtoma/test-speed/lib"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting...")

	// create an instance of our library
	speedService := speedTestLib.NewSpeedTest()

	// initialize the handlers
	speedService.InitHandlers()

	// start the server
	log.Fatal(http.ListenAndServe(":8888", nil))
}
