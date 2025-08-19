package form

type ForgotPassword struct {
	Email string `form:"email" validate:"required,email"`
}
