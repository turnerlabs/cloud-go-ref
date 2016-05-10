package app

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
//import ".././controllers/health"
//import ".././controllers/root"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from golang")
}

func Register() {

	r := mux.NewRouter()

	r.HandleFunc("/", Root)
	r.HandleFunc("/health", HealthCheck)

	http.Handle("/", r)

}
