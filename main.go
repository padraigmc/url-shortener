package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/padraigmc/url-shortener/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/link/shorten", handler.ShortenLink)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

	log.Info("Shorten URL started...")
	fmt.Println("Shorten URL started...")
}
