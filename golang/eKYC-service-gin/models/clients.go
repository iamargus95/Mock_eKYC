package models

import "github.com/jinzhu/gorm"

//Clients Structure.
type Client struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	Plan  Plan
}

//Plans Structure.
type Plan struct {
	gorm.Model
	ClientID uint
	Plan     string `gorm:"type:varchar(10)" json:"plan" validate:"required"`
}

//TableName return name of database table
func (d *Client) TableName() string {
	return "clients"
}

func (d *Plan) TableName() string {
	return "plans"
}
