package model

type ProfileRequest struct {
	Email string `json:"Email"`
}

type UpdateProfileRequest struct {
	Email    string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
