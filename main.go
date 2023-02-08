package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/padraigmc/url-shortener/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	rand.Seed(time.Now().UnixNano())
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/link/shorten", handler.ShortenLink)
	
	address := "localhost:10000"
	log.Info("Shorten URL starting on " + address)
	log.Fatal(http.ListenAndServe(address, myRouter))
}
