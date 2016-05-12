package main

import "net/http"

//import "github.com/turnerlabs/cloud-go-ref/routes"
//import "github.com/turnerlabs/cloud-go-ref/logger"
import "os"
import logrus "github.com/Sirupsen/logrus"

//   curl -i localhost:8080
//   curl -i localhost:8080/health

//var flarm string
var log logrus.Logger

type Controller struct {
}

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	log := logrus.WithFields(logrus.Fields{
		"app": "cloud-go-ref",
		"tbd": "to be determined common field...",
	})

	log.Info("cloud-go-ref is starting up...")

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Info("Attempting to listen on port: " + port)

	registerRoutes()

	log.Fatal(http.ListenAndServe(port, nil))

}
