package main

import "fmt"
import "net/http"
import "./routes/" 
import "os"

//   curl -i localhost:8080
//   curl -i localhost:8080/health

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":80"
	}

	routes.Register()
	fmt.Println("Listening on port: " + port)
	http.ListenAndServe(port, nil)

}
