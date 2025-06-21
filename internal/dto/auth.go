package dto

type SignUpInitRequest struct {
	Phone string `json:"phone" binding:"required,e164"`
}

type SignUpVerifyRequest struct {
	Phone string `json:"phone" binding:"required,e164"`
	Code  string `json:"code" binding:"required,min=6,max=6"`
}

type SignInRequest struct {
	Phone    string `json:"phone" binding:"required,e164"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}