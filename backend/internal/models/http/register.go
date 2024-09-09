package httpmodels

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Valid() (bool, string) {
	if len(r.Email) == 0 {
		return false, "email is empty"
	}
	if len(r.Password) == 0 {
		return false, "password is empty"
	}
	if !validateEmail(r.Email) {
		return false, "email is invalid"
	}
	return true, ""
}

type RegisterResponse struct {
	Registered bool `json:"registered"`
}
