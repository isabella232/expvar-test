package main

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var logger = log.WithFields(log.Fields{"app": os.Args[0]})

var images = []string {"gopher.png", "html.png", "java.png", "java.jpg", "javascript.png", "kotlin.png", "octocat.gif", "python.png", "scala.png"}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func main() {
	for {
		urls := []string {
			"http://localhost:8001/requests_per_minute",
			"http://localhost:8002/" + images[rand.Intn(len(images))],
			"http://localhost:8003/multiply",
		}
		for _, url := range urls {
			_, err := httpClient.Get(url)
			if err != nil {
				logger.Error(err)
			}
			//defer resp.Body.Close()
		}
		time.Sleep(time.Second * 1)
	}
}
