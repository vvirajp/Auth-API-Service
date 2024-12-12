package auth

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type User struct {
	Email string
}
