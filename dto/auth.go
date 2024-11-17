package dto

type SignUpRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`

	AccessToken string `json:"access_token"`
}
