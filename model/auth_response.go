package model

type RegisterResponse struct {
	Id          int64  `json:"-"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
}

type AuthResponse struct {
	Id          int64  `json:"-"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
}
