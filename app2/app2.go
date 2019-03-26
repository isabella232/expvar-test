package main

import (
	"expvar"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var logger = log.WithFields(log.Fields{"app": os.Args[0]})

var visits = expvar.NewInt("visits")
//var dummy = expvar.String{}

func main() {
	fs := http.FileServer(http.Dir("./app2/img"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visits.Add(1)
		fs.ServeHTTP(w, r)
	})
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		logger.Error(err)
	}
}
