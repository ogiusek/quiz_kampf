package dto

import (
	"fmt"
	usersconfig "lib/app/users/config"
	"lib/common/id"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenPayload struct {
	UserId id.ID
	Hash   string
}

func (refreshPayload RefreshTokenPayload) Encode() (RefreshToken, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = string(refreshPayload.UserId)
	claims["hash"] = string(refreshPayload.Hash)
	claims["exp"] = time.Now().Add(refreshDuration).Unix()

	unsignedToken := jwt.NewWithClaims(signingMethod, claims)
	token, err := unsignedToken.SignedString(usersconfig.Config.JwtSecret)
	if err != nil {
		return "", err
	}
	return RefreshToken(token), nil
}

func (tokenEncoded RefreshToken) DecodeSession() (RefreshTokenPayload, error) {
	token, err := jwt.Parse(string(tokenEncoded), func(token *jwt.Token) (interface{}, error) {
		return usersconfig.Config.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return RefreshTokenPayload{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return RefreshTokenPayload{}, fmt.Errorf("invalid claims")
	}
	session := RefreshTokenPayload{}
	sessionId, ok := claims["user_id"].(string)
	if !ok {
		return RefreshTokenPayload{}, fmt.Errorf("this is not a refresh token")
	}
	session.UserId = id.ID(sessionId)

	payloadHash, ok := claims["hash"].(string)
	if !ok {
		return RefreshTokenPayload{}, fmt.Errorf("this is not a refresh token")
	}
	session.Hash = payloadHash

	return session, nil
}
