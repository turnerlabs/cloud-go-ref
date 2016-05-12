package main

import (
	logrus "github.com/Sirupsen/logrus"
	"net/http"
)

func (rc Controller) root(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{"area": "rootController"}).Info("Root controller.")
	w.Write([]byte("Hello from cloud-go-ref"))
}
