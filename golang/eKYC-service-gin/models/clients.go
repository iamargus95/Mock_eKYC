package models

//Clients Structure.
type Client struct {
	Name  string `gorm:"type:varchar(50)" json:"name" validate:"required"`
	Email string `gorm:"type:varchar(50)" json:"email" validate:"required,email"`
	Plan  Plan
}

//Plans Structure.
type Plan struct {
	ClientID uint
	Plan     string `gorm:"type:varchar(10)" json:"plan" validate:"required"`
}
