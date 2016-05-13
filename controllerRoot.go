package main

import (
	logrus "github.com/Sirupsen/logrus"
	"net/http"
)

func (rc Controller) root(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{"area": "rootController"}).Info("Root controller.")
	w.Write([]byte("Hello from cloud-go-ref"))
}



func render() string {

  html := `<!doctype html>
           <head>
             <title>hello from golang!</title>
           </head>
           <body>
           <h1>Hello World!  This is cloud-go-ref...</h1>
           <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
           </body>
           </html>`

  return html

}
