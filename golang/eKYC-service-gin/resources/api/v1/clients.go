package v1resources

type Request struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Plan  string `binding:"required,oneof=basic advanced enterprise"`
}
