package models

import _ "github.com/jinzhu/gorm"

type Book struct {
	ID      uint   `json:"id" gorm:"primary_key;unique;not null"`
	Title   string `json:"title"`
	Release string `json:"release"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}
