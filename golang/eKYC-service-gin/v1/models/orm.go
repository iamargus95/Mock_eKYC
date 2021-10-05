package models

import "gorm.io/gorm"

//Structs for api/v1/signup
//Clients Structure.
type Client struct {
	gorm.Model
	Name      string `gorm:"unique_index"`
	Email     string `gorm:"unique_index"`
	Plan      Plan
	SecretKey SecretKey
	Pricing   Pricing
}

//Plans Structure.
type Plan struct {
	ID       uint `gorm:"primaryKey"` //Removing gorm.Model leads to duplicate field creation.
	ClientID uint
	Plan     string
	Pricing  Pricing
}

func (t *Client) TableName() string {
	return "clients"
}

func (t *Plan) TableName() string {
	return "plans"
}

//----------------------------------------------------------------------------------------------------------------------------------
//Structs for api/v1/image

type SecretKey struct {
	ID        uint `gorm:"primaryKey"`
	ClientID  uint
	Accesskey string
	Secretkey string
}

func (t *SecretKey) TableName() string {
	return "secretkey"
}

type Pricing struct {
	ClientID uint
	PlanID   uint
	Base     int
	Api_call int
	Storage  int
}

func (t *Pricing) TableName() string {
	return "pricing"
}
