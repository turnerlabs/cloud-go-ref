package main

import "fmt"
import "net/http"
import "github.com/turnerlabs/cloud-go-ref/routes"
import "os"

//   curl -i localhost:8080
//   curl -i localhost:8080/health

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	routes.Register()
	fmt.Println("Listening on port: " + port)
	fmt.Println(http.ListenAndServe(port, nil))

}
