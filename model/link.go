package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
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

// func DBMigrate(db *gorm.DB) *gorm.DB {
// 	db.AutoMigrate(&Link{})

// 	return db
// }