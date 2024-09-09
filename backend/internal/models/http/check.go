package httpmodels

type CheckTokenRequest struct {
	Token string `json:"token"`
}

func (r *CheckTokenRequest) Valid() (bool, string) {
	if len(r.Token) == 0 {
		return false, "token length is zero"
	}
	return true, ""
}

type CheckTokenResponse struct {
	Exists bool `json:"exists"`
}
