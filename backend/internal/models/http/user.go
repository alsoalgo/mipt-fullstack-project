package httpmodels

type EditProfileRequest struct {
	UserID    int64  `json:"userId,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SurName   string `json:"surName"`
}

func (r *EditProfileRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	if len(r.FirstName) <= 1 {
		return false, "firstName must have length greater than 1"
	}

	if len(r.LastName) <= 1 {
		return false, "lastName must have length greater than 1"
	}

	if len(r.SurName) <= 1 {
		return false, "surName must have length greater than 1"
	}

	return true, ""
}

type EditProfileResponse struct {
	Edited bool `json:"edited"`
}

type GetProfileRequest struct {
	UserID int64 `json:"userId,omitempty"`
}

func (r *GetProfileRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	return true, ""
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SurName   string `json:"surName"`
}

type GetProfileResponse struct {
	Info *User `json:"info"`
}
