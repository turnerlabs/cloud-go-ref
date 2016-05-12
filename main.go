package main

import "net/http"
import "github.com/turnerlabs/cloud-go-ref/routes"
import "github.com/turnerlabs/cloud-go-ref/logger"
import "os"
//import log "github.com/Sirupsen/logrus"

//   curl -i localhost:8080
//   curl -i localhost:8080/health

func main() {

  /*

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	logger := log.WithFields(log.Fields{
		"app": "cloud-go-ref",
		"tbd": "to be determined common field...",
	})

	logger.Info("cloud-go-ref is starting up...")

  */


	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	routes.Register()
	//logger.Info("Listening on port: " + port)
	//logger.Fatal(http.ListenAndServe(port, nil))
	http.ListenAndServe(port, nil)

}
