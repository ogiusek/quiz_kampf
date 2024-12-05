package dto

import (
	"encoding/json"
	"fmt"
	usersconfig "lib/app/users/config"
	"lib/app/users/models"
	"lib/common/id"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ogiusek/hw/src/hw"
)

var SessionHeader string = "Authorization"
var signingMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

type Session struct {
	UserId    id.ID            `json:"user_id"`
	UserName  models.UserName  `json:"user_name"`
	UserImage models.UserImage `json:"user_image"`
}

func ToSession(user models.User) Session {
	return Session{
		UserId:    user.UserId,
		UserName:  user.UserName,
		UserImage: user.UserImage,
	}
}

func (Session) GormDataType() string { return "jsonb" }
func (session *Session) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	json.Unmarshal(bytes, session)
	return nil
}

func (Session) FromHttp(r *http.Request) (any, hw.Resp) { return SessionFromHttp(r) }
func SessionFromHttp(r *http.Request) (Session, hw.Resp) {
	header := SessionToken(r.Header.Get("Authorization"))
	session, err := header.DecodeSession()
	if err != nil {
		resp := hw.NewResponse()
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(401)
		return Session{}, resp
	}
	return session, nil
}

func (session Session) Encode() (SessionToken, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = string(session.UserId)
	claims["user_name"] = string(session.UserName)
	claims["user_image"] = string(session.UserImage)
	claims["exp"] = time.Now().Add(sessionDuration).Unix()

	unsignedSessionToken := jwt.NewWithClaims(signingMethod, claims)
	sessionToken, err := unsignedSessionToken.SignedString(usersconfig.Config.JwtSecret)
	if err != nil {
		return "", err
	}

	return SessionToken(sessionToken), nil
}

func (tokenEncoded SessionToken) DecodeSession() (Session, error) {
	token, err := jwt.Parse(string(tokenEncoded), func(token *jwt.Token) (interface{}, error) {
		return usersconfig.Config.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return Session{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Session{}, fmt.Errorf("invalid claims")
	}

	session := Session{}

	sessionId, ok := claims["user_id"].(string)
	if !ok {
		return Session{}, fmt.Errorf("this is not a session token")
	}
	session.UserId = id.ID(sessionId)

	sessionUserName, ok := claims["user_name"].(string)
	if !ok {
		return Session{}, fmt.Errorf("this is not a session token")
	}
	session.UserName = models.UserName(sessionUserName)

	sessionUserImage, ok := claims["user_image"].(string)
	if !ok {
		return Session{}, fmt.Errorf("this is not a session token")
	}
	session.UserImage = models.UserImage(sessionUserImage)

	return session, nil
}

type Permission string

func (session Session) HasRight(permission Permission) error {
	return nil
}
