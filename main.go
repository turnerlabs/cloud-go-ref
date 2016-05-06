package main

import "fmt"
import "net/http"
import "./routes"
import "os"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	routes.Register()
	fmt.Println("Listening on port: " + port)
	http.ListenAndServe(port, nil)

}
