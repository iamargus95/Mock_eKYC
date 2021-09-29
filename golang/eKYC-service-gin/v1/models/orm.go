package models

import "gorm.io/gorm"

//Clients Structure.
type Client struct {
	gorm.Model
	Name  string `gorm:"unique_index"`
	Email string `gorm:"unique_index"`
	Plan  Plan
}

//Plans Structure.
type Plan struct {
	gorm.Model //Removing gorm.Model leads to duplicate field creation.
	ClientID   uint
	Plan       string
}

func (t *Client) TableName() string {
	return "clients"
}

func (t *Plan) TableName() string {
	return "plans"
}
