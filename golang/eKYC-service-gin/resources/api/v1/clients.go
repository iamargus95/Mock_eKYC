package v1resources

type ClientsResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty" validate:"required,email"`
	Plan  string `json:"required"`
}
