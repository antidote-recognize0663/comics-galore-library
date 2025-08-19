package form

type Signup struct {
	Email    string `form:"email" validate:"required,max=30,email"`
	Username string `form:"username" validate:"required,min=3,max=20"`
	Confirm  string `form:"confirm" validate:"required,confirm=Password"`
	Password string `form:"password" validate:"required,min=7,max=30,password"`
}
