package main

import (
	logrus "github.com/Sirupsen/logrus"
	"net/http"
	"html/template"
)

func (rc Controller) root(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{"area": "rootController"}).Info("Root controller.")
	//w.Write([]byte("Hello from cloud-go-ref"))


	html := render()

  //fmt.Println(html)

  tmpl, err := template.New("name").Parse(html)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

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
