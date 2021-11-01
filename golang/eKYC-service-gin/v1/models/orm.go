package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Structs for api/v1/signup
//Clients Structure.
type Client struct {
	gorm.Model
	Name      string `gorm:"unique_index"`
	Email     string `gorm:"unique_index"`
	Plan      Plan
	SecretKey SecretKey
}

//Plans Structure.
type Plan struct {
	ID       uint `gorm:"primaryKey"`
	ClientID uint
	Plan     string
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
}

func (t *SecretKey) TableName() string {
	return "secretkey"
}

type FileUpload struct {
	gorm.Model
	ClientID uint
	Type     string
	UUID     uuid.UUID
	Size     int64
}

func (t *FileUpload) TableName() string {
	return "file_upload"
}

type FaceMatch struct {
	gorm.Model
	ClientID uint
	Image1   uuid.UUID
	Image2   uuid.UUID
	Score    uint
}

func (t *FaceMatch) TableName() string {
	return "face_match"
}
