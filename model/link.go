package model

import (

)

type Link struct {
	Id				int
	Domain			string
	LongUrl			string
	ShortUrl		string
	ShortId			string
	// Clicks		int
	// UniqueClicks	int
}

func NewLink(domain string, longUrl string, shortUrl string, shortId string) *Link {
	url := Link{1, domain, longUrl, shortUrl, shortId}
	return &url
}