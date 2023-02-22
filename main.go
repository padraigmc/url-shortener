package main

import (
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/padraigmc/url-shortener/config"
	"github.com/padraigmc/url-shortener/handler"
	"github.com/padraigmc/url-shortener/model"
)

var db *gorm.DB

func main() {
	log.SetLevel(log.DebugLevel)
	rand.Seed(time.Now().UnixNano())

	isReady := &atomic.Value{}
	isReady.Store(false)

	// create database connection
	config := config.NewConfig()

	db, err := gorm.Open(mysql.Open(config.GetDbDSN()), &gorm.Config{})
   	if err != nil {
		log.Fatal("Could not connect database")
	}

	db.Migrator().DropTable(&model.Link{})
	db.Debug().AutoMigrate(&model.Link{})
	linkHandler := handler.LinkHandler{DB: db}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/link/shorten", linkHandler.ShortenLink)
	myRouter.HandleFunc("/link/{shortId:.*}", linkHandler.GetLinkFromShortUrl)
	myRouter.HandleFunc("/healthz", handler.Healthz)
	myRouter.HandleFunc("/readyz", handler.Readyz(isReady))
	
	address := config.Server.Host + ":" + config.Server.Port
	log.Info("Shorten URL starting on " + address)
	isReady.Store(true)
	log.Fatal(http.ListenAndServe(address, myRouter))
}