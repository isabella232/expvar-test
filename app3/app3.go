package main

import (
	"expvar"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var logger = log.WithFields(log.Fields{"app": os.Args[0]})

var visits = expvar.NewInt("visits")

func main() {
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		product := float64(1)
		for i := 0; i < rand.Intn(200); i++ {
			product *= rand.Float64() * math.Pow10(rand.Intn(50))
		}
		visits.Add(1)
		io.WriteString(w,  strconv.FormatFloat(product, 'f', -1, 64))
	})
	err := http.ListenAndServe(":8003", nil)
	if err != nil {
		logger.Error(err)
	}
}

