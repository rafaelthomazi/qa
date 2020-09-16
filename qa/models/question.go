package models

// Question represents a question made by one user
type Question struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Statement    string `json:"statement"`
	Answer       string `json:"answer"`
	AnswerUserID string `json:"answer_user_id"`
}
