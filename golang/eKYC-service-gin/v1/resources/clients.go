package v1resources

import "mime/multipart"

type SignupPayload struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Plan  string `binding:"required,oneof=basic advanced enterprise"`
}

type ImagePayload struct {
	Type string                `binding:"required,oneof=face id_card"`
	File *multipart.FileHeader `binding:"required"`
}
