package LoginCred

type LoginCredentials struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
