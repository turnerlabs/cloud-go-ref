package routes

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "github.com/turnerlabs/cloud-go-ref/pkg/controllers"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from golang")
}

func Register() {

	r := mux.NewRouter()

	r.HandleFunc("/", Root)
	r.HandleFunc("/health", HealthCheck)

	http.Handle("/", r)

}
