package main

import (
	"net/http"
)

//   curl -i localhost:8080/health

func (c Controller) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("App", "A reference app for cloud stuffs using golang")
	w.WriteHeader(200)
}
