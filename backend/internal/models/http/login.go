package httpmodels

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Valid() (bool, string) {
	if len(r.Email) == 0 {
		return false, "Email is empty"
	}

	if len(r.Password) == 0 {
		return false, "Password is empty"
	}

	if !validateEmail(r.Email) {
		return false, "Email is invalid"
	}

	return true, ""
}

func validateEmail(email string) bool {
	return true
}

type LoginResponse struct {
	Token string `json:"token"`
}
