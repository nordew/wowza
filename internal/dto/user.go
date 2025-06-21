package dto

type CreateUserRequest struct {
	ProfileName string `json:"profileName"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}