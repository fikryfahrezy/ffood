package model

type ProfileResponse struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type UpdateProfileResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
