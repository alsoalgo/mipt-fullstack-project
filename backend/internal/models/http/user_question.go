package httpmodels

type UserQuestion struct {
	Title    string `json:"title"`
	Question string `json:"question"`
}

type CreateQuestionRequest struct {
	UserID   int64  `json:"userId,omitempty"`
	Title    string `json:"title"`
	Question string `json:"question"`
}

func (r *CreateQuestionRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	if len(r.Title) < 10 {
		return false, "title is short"
	}
	if len(r.Question) < 10 {
		return false, "question is short"
	}

	return true, ""
}

type CreateQuestionResponse struct {
	Created bool `json:"created"`
}

type GetQuestionsRequest struct {
	UserID int64 `json:"userId,omitempty"`
}

func (r *GetQuestionsRequest) Valid() (bool, string) {
	if r.UserID <= 0 {
		return false, "userId must be positive"
	}

	return true, ""
}

type GetQuestionsResponse struct {
	Questions []*UserQuestion `json:"questions"`
}
