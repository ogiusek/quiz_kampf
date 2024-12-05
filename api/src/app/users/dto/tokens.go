package dto

type Tokens struct {
	SessionToken SessionToken `json:"session_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}
