package model

import (

)

type Link struct {
	Id				int		`json:"id,omitempty"`
	Domain			string	`json:"domain,omitempty"`
	Url				string	`json:"url,omitempty"`
	ShortUrl		string	`json:"short_url,omitempty"`
	ShortId			string	`json:"short_id,omitempty"`
}

func NewLink(domain string, longUrl string, shortUrl string, shortId string) *Link {
	url := Link{1, domain, longUrl, shortUrl, shortId}
	return &url
}