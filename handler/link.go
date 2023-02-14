package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"github.com/gorilla/mux"
	"github.com/padraigmc/url-shortener/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var urlAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type LinkHandler struct {
    DB *gorm.DB
}

func (h *LinkHandler) ShortenLink(w http.ResponseWriter, r *http.Request) {	
    log.Debug(fmt.Sprintf("%s request recieved to %s from %s", r.Method, r.RequestURI, r.RemoteAddr))

    // decode json in request body
    var link model.Link
    if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
        respondError(w, http.StatusBadRequest, err.Error())
		return
    }
    defer r.Body.Close()

    if !strings.HasPrefix(link.Url, "http://") && !strings.HasPrefix(link.Url, "https://") {
        link.Url = "https://" + link.Url
    }

    // generate url id
    parsedUrl, err := url.Parse(link.Url)
    if err != nil {
		log.Fatal(err)
	}

    link.Host = parsedUrl.Hostname()
    link.Path = parsedUrl.Path
    var shortId = generateShortId(5)
    link.ShortId = shortId
    link.ShortUrl = "short.en/" + shortId

    if err := h.DB.Save(&link).Error; err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
    log.Debug(fmt.Sprintf("URL '%s' mapped to %s", link.Url, link.ShortUrl))

    // write response
    respondJSON(w, http.StatusCreated, link)
}

// generate a string a length 'length' using aplphabet specified in 'urlAlphabet'
func generateShortId(length int) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = urlAlphabet[rand.Intn(len(urlAlphabet))]
    }
    return string(b)
}

func (h *LinkHandler) GetLinkFromShortUrl(w http.ResponseWriter, r *http.Request) {
	link := model.Link{}
	vars := mux.Vars(r)
	shortId := vars["shortId"]

	// if short Url supplied, extract Id
	if split := strings.Split(shortId, "/") ; len(split) > 1 {
		shortId = split[len(split)-1]
	}

	if err := h.DB.First(&link, model.Link{ShortId: shortId}).Error; err != nil {
        respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	respondJSON(w, http.StatusOK, link)
}