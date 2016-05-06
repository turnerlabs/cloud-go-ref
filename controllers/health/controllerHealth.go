package controllerHealth

import (
	"net/http"
)

//   curl -i localhost:8080/health

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("App", "A reference app for cloud stuffs using golang")
	w.WriteHeader(200)

}
