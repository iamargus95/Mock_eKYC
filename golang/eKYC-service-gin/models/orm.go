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
	ClientID uint   `sql:"index"`
	Plan     string `gorm:"one2one:plan_clients"`
}

func (t *Client) TableName() string {
	return "clients"
}

func (t *Plan) TableName() string {
	return "plans"
}
