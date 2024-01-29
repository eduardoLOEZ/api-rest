package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	USERNAME string `json:"username" gorm:"text;not null;default:null`

	EMAIL   string `json:"email" gorm:"text;not null;default:null`

}
