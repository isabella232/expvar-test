package main

import (
	"expvar"
	"github.com/paulbellamy/ratecounter"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var logger = log.WithFields(log.Fields{"app": os.Args[0]})

var (
	visits = expvar.NewInt("visits")
	counter *ratecounter.RateCounter
	requestsPerMinute = expvar.NewInt("requests_per_minute")
)

func increment(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	counter.Incr(1)
	requestsPerMinute.Set(counter.Rate())
	io.WriteString(w,  strconv.FormatInt(counter.Rate(), 10))
}

func main() {
	counter = ratecounter.NewRateCounter(1 * time.Minute)
	http.HandleFunc("/requests_per_minute", increment)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		logger.Error(err)
	}
}
