package model

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

type Link struct {
	gorm.Model
	Host			string		`json:"host,omitempty"`
	Url				string		`gorm:"not null;size:256" json:"url,omitempty"`
	Path			string		`json:"path,omitempty"`
	ShortUrl		string		`gorm:"not null" json:"short_url,omitempty"`
	ShortId			string		`gorm:"not null" json:"short_id,omitempty"`
	Clicks			int			`gorm:"default:0;not null" json:"clicks,omitempty"`
}