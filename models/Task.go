package models

import "gorm.io/gorm"

type Task struct{

	gorm.Model

	Title 		string  `gorm:"type:varchar(100);not null; unique_index" json:"title"`
	Description string	`json:"description"`
	Dond 		bool 	`gorm:"default:false"`
	UserID 		uint	`json:"user_id"`
}