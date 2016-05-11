package routes

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "github.com/turnerlabs/cloud-go-ref/controllers"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from golang")
}

func Register() {

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Root)
	r.HandleFunc("/health", controllers.HealthCheck)

	http.Handle("/", r)

}
