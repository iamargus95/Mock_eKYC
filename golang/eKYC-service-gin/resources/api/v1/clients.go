package v1resources

type Request struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Plan  string `json:"plan," validate:"required,plan"`
}
