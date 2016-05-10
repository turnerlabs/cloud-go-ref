package app

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from cloud-go-ref"))
}
