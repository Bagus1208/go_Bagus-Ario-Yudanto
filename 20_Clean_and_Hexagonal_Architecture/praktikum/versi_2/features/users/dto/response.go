package dto

type LoginResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"token"`
}
