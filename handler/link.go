package handler

import (
    "net/http"
    "math/rand"
    // "encoding/json"
    "io/ioutil"
    log "github.com/sirupsen/logrus"
    // "github.com/padraigmc/url-shortener/model"
)

func ShortenLink(w http.ResponseWriter, r *http.Request) {
    log.Debug("/link/shorted")
	var shortId = generateShortId(5)
    var shortUrl = "short.en/" + shortId
    requestBody, _ := ioutil.ReadAll(r.Body)
    log.Debug(requestBody)
    log.Debug(shortUrl)
    
    // link := model.NewLink(r.URL.Host, r.URL.RequestURI(), shortUrl, shortId)

    // json.Marshal(link)
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateShortId(length int) string {
    b := make([]rune, length)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}