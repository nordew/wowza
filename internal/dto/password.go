package dto

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordConfirmRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ResetPasswordConfirmCompleteRequest struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}