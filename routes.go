package main

import "net/http"
import "github.com/gorilla/mux"
import logrus "github.com/Sirupsen/logrus"

func registerRoutes() {

	log.WithFields(logrus.Fields{"area": "registerRoutes"}).Info("Registering routes.")

	r := mux.NewRouter()

	c := Controller{}

	r.HandleFunc("/", c.root)
	r.HandleFunc("/health", c.health)

	http.Handle("/", r)

}
