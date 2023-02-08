package handler

import (
    "fmt"
    "net/http"
    "math/rand"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    "github.com/padraigmc/url-shortener/model"
)

var urlAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortenLink(w http.ResponseWriter, r *http.Request) {	
    log.Debug(fmt.Sprintf("%s request recieved to %s from %s", r.Method, r.RequestURI, r.RemoteAddr))

    // decode json in request body
    var link model.Link
    err := json.NewDecoder(r.Body).Decode(&link)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // generate url id
    var shortId = generateShortId(5)
    link.ShortId = shortId
    link.ShortUrl = "short.en/" + shortId

    log.Debug(fmt.Sprintf("URL '%s' mapped to %s", link.Url, link.ShortUrl))

    // write response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(link)
}

// generate a string a length 'length' using aplphabet specified in 'urlAlphabet'
func generateShortId(length int) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = urlAlphabet[rand.Intn(len(urlAlphabet))]
    }
    return string(b)
}