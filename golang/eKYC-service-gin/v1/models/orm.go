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
	gorm.Model //Removing gorm.Model leads to duplicate field creation.
	ClientID   uint
	Plan       string
}

type MiddleMap struct {
	gorm.Model
	ClientID uint
	PlanID   uint
}

func (t *Client) TableName() string {
	return "clients"
}

func (t *Plan) TableName() string {
	return "plans"
}

func (t *MiddleMap) TableName() string {
	return "middle_map"
}
