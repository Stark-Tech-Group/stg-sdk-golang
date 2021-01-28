package response

type AuthResponse struct {
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
	TokenType    string   `json:"token_type"`
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int      `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
}
