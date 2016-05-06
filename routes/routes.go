package routes

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import ".././controllers/health"
import ".././controllers/root"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there from golang")
}

//   curl -i localhost:8080/health

func Register() {

  r := mux.NewRouter()

  r.HandleFunc("/", controllerRoot.Root)
  r.HandleFunc("/health", controllerHealth.HealthCheck)

  http.Handle("/", r)

}
