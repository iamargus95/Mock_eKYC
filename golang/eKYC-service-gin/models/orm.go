package models

import "gorm.io/gorm"

//Clients Structure.
type Client struct {
	gorm.Model
	Name  string
	Email string
	Plan  Plan
}

//Plans Structure.
type Plan struct {
	gorm.Model
	ClientID uint
	Plan     string
}
